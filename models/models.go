package models

import (
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"go.mongodb.org/mongo-driver/mongo/address"
)

type UserData struct {
	AssetAAmount   *big.Float
	AssetBAmount   *big.Float
	AssetAPriceUSD *big.Float
	AssetBPriceUSD *big.Float
}

type ProtocolData struct {
	Name        string
	TotalSupply *big.Int
	Reserves    struct {
		Reserve0           *big.Int
		Reserve1           *big.Int
		BlockTimestampLast uint32
	}
}

type LMCAddress struct {
	LmcName string
	Address common.Address
}

type Token struct {
	Address map[string]address.Address `json:"address" bson:"address"`
	Symbol  string                     `json:"symbol" bson:"symbol"`
	Name    string                     `json:"name" bson:"name"`

	AmountBN  AmountBigFloat `json:"amountBN" bson:"amountBN"`
	AmountUSD AmountBigFloat `json:"amountUSD" bson:"amountUSD"`
	PriceUSD  AmountBigFloat `json:"priceUSD" bson:"priceUSD"`
}

type AmountBigFloat struct {
	F string `json:"f" bson:"f"`
}

type UserPosition struct {
	AssetAAmount   *big.Float
	AssetBAmount   *big.Float
	AssetAPriceUSD *big.Float
	AssetBPriceUSD *big.Float
}

type LMCampaign struct {
	Address               common.Address
	TotalStaked           *big.Int
	AssetAPortion         *big.Int
	AssetBPortion         *big.Int
	Tokens                []*Token
	CurrentRate           *big.Float
	AssetAPortionValueUSD *big.Float
	AssetBPortionValueUSD *big.Float
}

type Price struct {
	currency string
	value    int
}

type TokenPrice struct {
	tokenId string
	price   Price
}

type TenantData struct {
	Address  common.Address
	Protocol string
}

type Config struct {
	FE ConfigFE `json:"fe" bson:"fe"`
}

type UniswapObject struct {
	RewardContracts map[string]interface{} `json:"rewardContracts" bson:"rewardContracts"`
}

type ConfigFE struct {
	Uniswap UniswapObject `json:"uniswap" bson:"uniswap"`
}

type Tenant struct {
	DisplayName string `json:"displayName" bson:"displayName"`
	Config      Config `json:"config" bson:"config"`
}

type Pool struct {
	Protocol        string            `json:"protocol" bson:"protocol"`
	StakerAddr      address.Address   `json:"staker" bson:"staker"`
	RTokensAddrs    []address.Address `json:"rewardTokens" bson:"rewardTokens"`
	BonusTokenAddrs address.Address   `json:"bonusToken" bson:"bonusToken"`
	LPAddr          address.Address   `json:"liquidityPool" bson:"liquidityPool"`
	LPTokenAddr     address.Address   `json:"liquidityPoolToken" bson:"liquidityPoolToken"`
	PTokensAddrs    []address.Address `json:"provisionTokens" bson:"provisionTokens"`
}

type StakingConfig struct {
	Pools []*Pool `json:"pools" bson:"pools"`
}

type Action struct {
	Name          string                 `json:"name" bson:"name"`
	LPTokenAmount map[string]interface{} `json:"lpTokenAmount" bson:"lpTokenAmount"`
}

type LPAssets struct {
	LPToken           *Token                `json:"liquidityToken" bson:"liquidityToken"`
	ReserveTokens     []*Token              `json:"reserveTokens" bson:"reserveTokens"`
	LPTokenToReserves map[string]*big.Float `json:"lpTokenToReserves" bson:"lpTokenToReserves"`
}

type TxData struct {
	// Hash   address.Hash    `json:"hash" bson:"hash"`
	// From   address.Address `json:"from" bson:"from"`
	// To     address.Address `json:"to" bson:"to"`
	Status uint64 `json:"status" bson:"status"`
	Data   string `json:"data" bson:"data"`
	// CumulativeGasUsed uint64          `json:"cumulativeGasUsed" bson:"cumulativeGasUsed"`
	GasUsed  uint64 `json:"gasUsed" bson:"gasUsed"`
	GasLimit uint64 `json:"gasLimit" bson:"gasLimit"`
	GasPrice uint64 `json:"gasPrice" bson:"gasPrice"`
	// NativeTokenPrice *bigfloat.Float `json:"ethPrice" bson:"ethPrice"`
	Value       uint64    `json:"value" bson:"value"`
	Nonce       uint64    `json:"nonce" bson:"nonce"`
	BlockTime   time.Time `json:"blockTime" bson:"blockTime"`
	BlockNumber uint64    `json:"blockNumber" bson:"blockNumber"`

	Protocol string `json:"protocol" bson:"protocol"`
	// LiquidityPool address.Address `json:"liquidityPool" bson:"liquidityPool"`
	Timestamp time.Time `json:"timestamp" bson:"timestamp"`
	Action    *Action   `json:"action" bson:"action"`
	LPAssets  *LPAssets `json:"liquidityPoolAssets" bson:"liquidityPoolAssets"`
}
