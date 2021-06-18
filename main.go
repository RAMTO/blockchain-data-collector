package main

import (
	"blockchain-data-collector/formulas"
	"blockchain-data-collector/lmc"
	"blockchain-data-collector/models"
	simplelogger "blockchain-data-collector/simple-logger"
	"blockchain-data-collector/uniswapPair"
	"context"
	"fmt"
	"log"
	"math/big"
	"os"
	"reflect"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func itemExists(arrayType interface{}, item interface{}) bool {
	arr := reflect.ValueOf(arrayType)

	if arr.Kind() != reflect.Slice {
		panic("Invalid data-type")
	}

	for i := 0; i < arr.Len(); i++ {
		if arr.Index(i).Interface() == item {
			return true
		}
	}

	return false
}

func addressExists(array []*common.Address, item *common.Address) bool {
	for i := 0; i < len(array); i++ {
		if reflect.DeepEqual(array[i], item) {
			return true
		}
	}

	return false
}

func initAndConnectDB() (*mongo.Client, error) {
	// Get env file
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
		return nil, err
	}

	// Client init
	clientOptions := options.Client().ApplyURI(os.Getenv("MONGO_DB"))

	client, err := mongo.NewClient(clientOptions)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	err = client.Connect(context.Background())
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return client, nil
}

func initAndConnectEtherium() (*ethclient.Client, error) {
	client, err := ethclient.Dial("http://ethereumnode.defiterm-dev.net:8545")
	if err != nil {
		return nil, err
	}

	return client, nil
}

func getLMCAddresses(collection *mongo.Collection) []*models.LMCAddress {
	addresses := []*models.LMCAddress{}

	cursor, err := collection.Find(context.TODO(), bson.M{"name": "dobreff"})
	if err != nil {
		log.Fatal(err)
	}

	for cursor.Next(context.TODO()) {
		var data models.Tenant
		err := cursor.Decode(&data)

		if err != nil {
			log.Fatal(err)
		}

		if len(data.Config.FE.Uniswap.RewardContracts) > 0 {
			for key, pair := range data.Config.FE.Uniswap.RewardContracts {
				pairAddress := common.Address{}

				// type assertion for interfaces
				if p, ok := pair.(string); ok {
					pairAddress = common.HexToAddress(p)

					singleAddress := models.LMCAddress{
						LmcName: key,
						Address: common.Address(pairAddress),
					}

					addresses = append(addresses, &singleAddress)
				}
			}
		}
	}

	if err := cursor.Err(); err != nil {
		log.Fatal(err)
	}

	return addresses
}

func getTxData(collection *mongo.Collection, addresses []*models.LMCAddress) []*models.UserPosition {
	positions := []*models.UserPosition{}
	lmcAddresses := []*common.Address{}

	for _, singleAddress := range addresses {
		lmcAddresses = append(lmcAddresses, &singleAddress.Address)
	}

	cursor, err := collection.Find(context.TODO(), bson.M{})
	if err != nil {
		log.Fatal(err)
	}

	for cursor.Next(context.TODO()) {
		var data models.TxData
		err := cursor.Decode(&data)

		addressPointer := common.HexToAddress(data.LiquidityPool.A)

		if addressExists(lmcAddresses, &addressPointer) && len(data.LPAssets.ReserveTokens) > 0 {
			assetAAmount, _ := new(big.Float).SetString(data.LPAssets.ReserveTokens[0].AmountBN.F)
			assetBAmount, _ := new(big.Float).SetString(data.LPAssets.ReserveTokens[1].AmountBN.F)
			assetAPriceUSD, _ := new(big.Float).SetString(data.LPAssets.ReserveTokens[0].AmountUSD.F)
			assetBPriceUSD, _ := new(big.Float).SetString(data.LPAssets.ReserveTokens[0].AmountUSD.F)

			position := models.UserPosition{
				AssetAAmount:   assetAAmount,
				AssetBAmount:   assetBAmount,
				AssetAPriceUSD: assetAPriceUSD,
				AssetBPriceUSD: assetBPriceUSD,
			}

			if err != nil {
				log.Fatal(err)
			}

			positions = append(positions, &position)
		}
	}

	if err := cursor.Err(); err != nil {
		log.Fatal(err)
	}

	return positions
}

func getCampaignsData(client *ethclient.Client, addresses []*models.LMCAddress, poolData *models.ProtocolData) []*models.LMCampaign {
	campaigns := []*models.LMCampaign{}

	for _, address := range addresses {
		lmcContract, err := lmc.NewLmc(address.Address, client)
		if err != nil {
			log.Fatal(err)
		}

		// Get total staked
		totalStaked, err := lmcContract.LmcCaller.TotalStaked(&bind.CallOpts{})
		if err != nil {
			log.Fatal(err)
		}

		// Get portion A tokens
		assetAPortion := new(big.Int).Div(
			new(big.Int).Mul(
				totalStaked,
				poolData.Reserves.Reserve0,
			),
			poolData.TotalSupply,
		)

		// Get portion B tokens
		assetBPortion := new(big.Int).Div(
			new(big.Int).Mul(
				totalStaked,
				poolData.Reserves.Reserve1,
			),
			poolData.TotalSupply,
		)

		priceAUSD := new(big.Float).SetFloat64(0.39)
		priceBUSD := new(big.Float).SetFloat64(2400)

		// Get decimals from Tokens
		decimals := new(big.Float).SetInt(big.NewInt(int64(1000000000000000000)))

		assetAPortionEthers := new(big.Float).Quo(
			new(big.Float).SetInt(assetAPortion),
			decimals,
		)

		assetBPortionEthers := new(big.Float).Quo(
			new(big.Float).SetInt(assetBPortion),
			decimals,
		)

		assetAPortionValueUSD := new(big.Float).Mul(
			assetAPortionEthers,
			priceAUSD,
		)

		assetBPortionValueUSD := new(big.Float).Mul(
			assetBPortionEthers,
			priceBUSD,
		)

		currentRate := big.NewFloat(float64(0))

		// Get current rate
		if assetAPortion.Cmp(big.NewInt(int64(0))) != 0 && assetBPortion.Cmp(big.NewInt(int64(0))) != 0 {
			currentRate = new(big.Float).Quo(
				new(big.Float).SetInt(assetAPortion),
				new(big.Float).SetInt(assetBPortion),
			)
		}

		campaign := models.LMCampaign{
			Address:               address.Address,
			TotalStaked:           totalStaked,
			CurrentRate:           currentRate,
			AssetAPortion:         assetAPortion,
			AssetBPortion:         assetBPortion,
			AssetAPortionPriceUSD: priceAUSD,
			AssetBPortionPriceUSD: priceBUSD,
			AssetAPortionValueUSD: assetAPortionValueUSD,
			AssetBPortionValueUSD: assetBPortionValueUSD,
		}

		campaigns = append(campaigns, &campaign)
	}

	return campaigns
}

func getPoolPairData(client *ethclient.Client, address common.Address) *models.ProtocolData {
	uniswapPairContract, err := uniswapPair.NewUniswapPair(address, client)
	if err != nil {
		log.Fatal(err)
	}

	reserves, err := uniswapPairContract.UniswapPairCaller.GetReserves(&bind.CallOpts{})
	if err != nil {
		log.Fatal(err)
	}

	totalSuply, err := uniswapPairContract.UniswapPairCaller.TotalSupply(&bind.CallOpts{})
	if err != nil {
		log.Fatal(err)
	}

	protocolData := models.ProtocolData{
		Name:        "Uniswap",
		TotalSupply: totalSuply,
		Reserves:    reserves,
	}

	return &protocolData
}

func main() {
	fmt.Println("")
	fmt.Println("=================")
	fmt.Println("=== APP START ===")

	// Init and connect to MongoDB
	clientDB, err := initAndConnectDB()
	if err != nil {
		log.Fatal(err)
	}

	dbTenant := clientDB.Database("LMaaS-TenantData")
	collectionProjects := dbTenant.Collection("projects")

	dbClients := clientDB.Database("LMaaS-ClientData")
	collectionTransactions := dbClients.Collection("txdata_dobreff_rinkeby")

	// Init etherium client
	client, err := initAndConnectEtherium()
	if err != nil {
		log.Fatal(err)
	}

	// Get Uniswap data
	uniswapPoolData := getPoolPairData(client, common.HexToAddress("0x0Bd2f8af9f5E5BE43B0DA90FE00A817e538B9306"))

	// Get all tenant LMC addresses
	addresses := getLMCAddresses(collectionProjects)

	// Get transactions data
	positions := getTxData(collectionTransactions, addresses)

	// Get campaigns data
	campaigns := getCampaignsData(client, addresses, uniswapPoolData)

	// for _, position := range positions {
	// 	spew.Dump(position)
	// }

	for _, singleLMC := range campaigns {
		simplelogger.Log("LMC: ", singleLMC.Address)

		// Check for total staked
		if singleLMC.TotalStaked.Cmp(big.NewInt(int64(0))) != 0 {

			// Calculate current LMC assets value at current price
			currentAssetsCurrentPrices := formulas.CalculateAssetsValue(
				[]*big.Float{singleLMC.AssetAPortionValueUSD, singleLMC.AssetBPortionValueUSD},
			)

			// Inital values
			poolTotalILUSD := new(big.Float).SetFloat64(0)
			totalInitialAssetsCurrentPrices := new(big.Float).SetFloat64(0)

			for _, position := range positions {
				// Exchange rates of postion
				exchangeRatePosition := new(big.Float).Quo(
					position.AssetAAmount,
					position.AssetBAmount,
				)

				// Exchange rates of pool
				exchangeRateCurrent := new(big.Float).Quo(
					new(big.Float).SetInt(singleLMC.AssetAPortion),
					new(big.Float).SetInt(singleLMC.AssetBPortion),
				)

				r := new(big.Float).Quo(exchangeRateCurrent, exchangeRatePosition)

				// Calculate initial assets at initial price
				assetAInitialPrice := formulas.CalculateAssetPrice(
					position.AssetAAmount,
					position.AssetAPriceUSD,
				)

				assetBInitialPrice := formulas.CalculateAssetPrice(
					position.AssetBAmount,
					position.AssetBPriceUSD,
				)

				initialAssetsInitialPrices := formulas.CalculateAssetsValue(
					[]*big.Float{assetAInitialPrice, assetBInitialPrice},
				)

				// Calculate initial assets at current price
				initialAssetACurrentPrice := formulas.CalculateAssetPrice(
					position.AssetAAmount,
					singleLMC.AssetAPortionPriceUSD,
				)

				initialAssetBCurrentPrice := formulas.CalculateAssetPrice(
					position.AssetBAmount,
					singleLMC.AssetBPortionPriceUSD,
				)

				initialAssetsCurrentPrices := formulas.CalculateAssetsValue(
					[]*big.Float{initialAssetACurrentPrice, initialAssetBCurrentPrice},
				)

				// Needed for pool trading fees
				totalInitialAssetsCurrentPrices = new(big.Float).Add(
					totalInitialAssetsCurrentPrices,
					initialAssetsCurrentPrices,
				)

				il := formulas.CalculateImpermanentLoss(r)
				ilUSD := formulas.CalculateImpermanentLossUSD(initialAssetsInitialPrices, il)

				poolTotalILUSD = new(big.Float).Add(poolTotalILUSD, ilUSD)
			}

			poolTradingFees := formulas.CalculateFees(
				currentAssetsCurrentPrices,
				totalInitialAssetsCurrentPrices,
				poolTotalILUSD,
			)

			poolNetGainLoss := formulas.CalculateNetGainLoss(
				totalInitialAssetsCurrentPrices,
				currentAssetsCurrentPrices,
			)

			fmt.Printf("%+v\n", "")
			fmt.Printf("%+v\n", "Pool data:")
			fmt.Printf("%+v\n", "* Trading fees:")
			fmt.Printf("%+v\n", poolTradingFees)
			fmt.Printf("%+v\n", "* IL USD:")
			fmt.Printf("%+v\n", poolTotalILUSD)
			fmt.Printf("%+v\n", "* Net Gain/Loss:")
			fmt.Printf("%+v\n", poolNetGainLoss)
		}
	}
}