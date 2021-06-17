package main

import (
	impermanentloss "blockchain-data-collector/formulas/impermanent-loss"
	netgainloss "blockchain-data-collector/formulas/net-gain-loss"
	tradingfees "blockchain-data-collector/formulas/trading-fees"
	"blockchain-data-collector/lmc"
	"blockchain-data-collector/models"
	"blockchain-data-collector/uniswapPair"
	"context"
	"fmt"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func initAndConnectDB() *mongo.Client {
	// Client init
	clientOptions := options.Client().ApplyURI("")

	client, err := mongo.NewClient(clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	err = client.Connect(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	return client
}

func initAndConnectEtherium() *ethclient.Client {
	client, err := ethclient.Dial("http://ethereumnode.defiterm-dev.net:8545")
	if err != nil {
		fmt.Println(err)
	}

	return client
}

func getLMCAddresses(collection *mongo.Collection) []*models.LMCAddress {
	addresses := []*models.LMCAddress{}

	cursor, err := collection.Find(context.TODO(), bson.M{})
	if err != nil {
		log.Fatal(err)
	}

	for cursor.Next(context.TODO()) {
		var data models.Tenant
		err := cursor.Decode(&data)

		if len(data.Config.FE.Uniswap.RewardContracts) > 0 {
			for key, pair := range data.Config.FE.Uniswap.RewardContracts {
				// type assertion for interfaces
				pairAddress := common.HexToAddress(pair.(string))

				singleAddress := models.LMCAddress{
					LmcName: key,
					Address: common.Address(pairAddress),
				}

				addresses = append(addresses, &singleAddress)
			}
		}

		if err != nil {
			log.Fatal(err)
		}
	}

	if err := cursor.Err(); err != nil {
		log.Fatal(err)
	}

	return addresses
}

func getTxData(collection *mongo.Collection) []*models.UserPosition {
	positions := []*models.UserPosition{}

	cursor, err := collection.Find(context.TODO(), bson.M{})
	if err != nil {
		log.Fatal(err)
	}

	for cursor.Next(context.TODO()) {
		var data models.TxData
		err := cursor.Decode(&data)

		assetAAmount, ok := new(big.Float).SetString(data.LPAssets.ReserveTokens[0].AmountBN.F)
		_ = ok
		assetBAmount, ok := new(big.Float).SetString(data.LPAssets.ReserveTokens[1].AmountBN.F)
		_ = ok
		assetAPriceUSD, ok := new(big.Float).SetString(data.LPAssets.ReserveTokens[0].AmountUSD.F)
		_ = ok
		assetBPriceUSD, ok := new(big.Float).SetString(data.LPAssets.ReserveTokens[0].AmountUSD.F)
		_ = ok

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
			AssetAPortion:         assetAPortion,
			AssetBPortion:         assetBPortion,
			CurrentRate:           currentRate,
			AssetAPortionValueUSD: assetAPortionValueUSD,
			AssetBPortionValueUSD: assetBPortionValueUSD,
		}

		campaigns = append(campaigns, &campaign)
	}

	return campaigns
}

func getUniswapData(client *ethclient.Client, address common.Address) *models.ProtocolData {
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
	// Init and connect to MongoDB
	clientDB := initAndConnectDB()

	dbTenant := clientDB.Database("LMaaS-TenantData")
	collectionProjects := dbTenant.Collection("projects")

	dbClients := clientDB.Database("LMaaS-ClientData")
	collectionTransactions := dbClients.Collection("txdata_dobreff_rinkeby")

	// Init etherium client
	client := initAndConnectEtherium()

	// Get Uniswap data
	uniswapPoolData := getUniswapData(client, common.HexToAddress("0x0Bd2f8af9f5E5BE43B0DA90FE00A817e538B9306"))

	// Get all tenant LMC addresses
	addresses := getLMCAddresses(collectionProjects)

	// Get transactions data
	positions := getTxData(collectionTransactions)

	_ = positions

	// Get campaigns data
	campaigns := getCampaignsData(client, addresses, uniswapPoolData)

	_ = campaigns

	// for _, singleLMC := range campaigns {
	// 	spew.Dump(singleLMC)
	// }

	// for _, position := range positions {
	// spew.Dump(position)
	// }

	// Pool current data
	priceMapping := make(map[string]*big.Float)
	priceMapping["assetACurrent"] = new(big.Float).SetInt(campaigns[1].AssetAPortion)
	priceMapping["assetACurrentUSD"] = campaigns[1].AssetAPortionValueUSD

	priceMapping["assetBCurrent"] = new(big.Float).SetInt(campaigns[1].AssetBPortion)
	priceMapping["assetBCurrentUSD"] = campaigns[1].AssetBPortionValueUSD

	// Calculate current assets at current price
	assetACurrentPrice := tradingfees.CalculateAssetPrice(
		priceMapping["assetACurrent"],
		priceMapping["assetACurrentUSD"],
	)

	assetBCurrentPrice := tradingfees.CalculateAssetPrice(
		priceMapping["assetBCurrent"],
		priceMapping["assetBCurrentUSD"],
	)

	currentAssetsCurrentPrices := tradingfees.CalculateAssetsValue(
		[]*big.Float{assetACurrentPrice, assetBCurrentPrice},
	)

	// Inital values
	poolTotalILUSD := new(big.Float).SetFloat64(0)
	totalInitialAssetsCurrentPrices := new(big.Float).SetFloat64(0)

	for _, position := range positions {
		// Exchange rates
		exchangeRateInitial := new(big.Float).Quo(position.AssetAAmount, position.AssetBAmount)
		exchangeRateCurrent := new(big.Float).Quo(priceMapping["assetACurrent"], priceMapping["assetBCurrent"])

		r := new(big.Float).Quo(exchangeRateCurrent, exchangeRateInitial)

		// Calculate initial assets at initial price
		assetAInitialPrice := tradingfees.CalculateAssetPrice(
			position.AssetAAmount,
			position.AssetAPriceUSD,
		)

		assetBInitialPrice := tradingfees.CalculateAssetPrice(
			position.AssetBAmount,
			position.AssetBPriceUSD,
		)

		initialAssetsInitialPrices := tradingfees.CalculateAssetsValue(
			[]*big.Float{assetAInitialPrice, assetBInitialPrice},
		)

		// Calculate initial assets at current price
		initialAssetACurrentPrice := tradingfees.CalculateAssetPrice(
			position.AssetAAmount,
			priceMapping["assetACurrentUSD"],
		)

		initialAssetBCurrentPrice := tradingfees.CalculateAssetPrice(
			position.AssetBAmount,
			priceMapping["assetBCurrentUSD"],
		)

		initialAssetsCurrentPrices := tradingfees.CalculateAssetsValue(
			[]*big.Float{initialAssetACurrentPrice, initialAssetBCurrentPrice},
		)

		// Needed for pool trading fees
		totalInitialAssetsCurrentPrices = new(big.Float).Add(totalInitialAssetsCurrentPrices, initialAssetsCurrentPrices)

		il := impermanentloss.CalculateImpermanentLoss(r)
		ilUSD := impermanentloss.CalculateImpermanentLossUSD(initialAssetsInitialPrices, il)

		poolTotalILUSD = new(big.Float).Add(poolTotalILUSD, ilUSD)
	}

	poolTradingFees := tradingfees.CalculateFees(
		currentAssetsCurrentPrices,
		totalInitialAssetsCurrentPrices,
		poolTotalILUSD,
	)

	poolNetGainLoss := netgainloss.CalculateNetGainLoss(
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
