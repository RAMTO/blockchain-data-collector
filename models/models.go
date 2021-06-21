package models

import (
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/common"
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
	Tokens      []*TokenData
	Reserves    struct {
		Reserve0           *big.Int
		Reserve1           *big.Int
		BlockTimestampLast uint32
	}
}

type TokenData struct {
	Name     string
	Address  common.Address
	Decimals int
}

type LMCAddress struct {
	LmcName string
	Address common.Address
}

type Token struct {
	Address map[string]string `json:"address" bson:"address"`
	Symbol  string            `json:"symbol" bson:"symbol"`
	Name    string            `json:"name" bson:"name"`

	AmountBN  AmountBigFloat `json:"amountBN" bson:"amountBN"`
	AmountUSD AmountBigFloat `json:"amountUSD" bson:"amountUSD"`
	PriceUSD  AmountBigFloat `json:"priceUSD" bson:"priceUSD"`
}

type AmountBigFloat struct {
	F string `json:"f" bson:"f"`
}

type UserPosition struct {
	ExchangeRate   *big.Float
	AssetAAmount   *big.Float
	AssetBAmount   *big.Float
	AssetAPriceUSD *big.Float
	AssetBPriceUSD *big.Float
	AssetAValueUSD *big.Float
	AssetBValueUSD *big.Float
	AssetsValueUSD *big.Float
}

type LMCampaign struct {
	Address               common.Address
	TotalStaked           *big.Int
	Tokens                []*Token
	CurrentRate           *big.Float
	AssetAPortion         *big.Int
	AssetBPortion         *big.Int
	AssetAPortionPriceUSD *big.Float
	AssetBPortionPriceUSD *big.Float
	AssetAPortionValueUSD *big.Float
	AssetBPortionValueUSD *big.Float
	AssetsValueUSD        *big.Float
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
	Protocol        string   `json:"protocol" bson:"protocol"`
	StakerAddr      string   `json:"staker" bson:"staker"`
	RTokensAddrs    []string `json:"rewardTokens" bson:"rewardTokens"`
	BonusTokenAddrs string   `json:"bonusToken" bson:"bonusToken"`
	LPAddr          string   `json:"liquidityPool" bson:"liquidityPool"`
	LPTokenAddr     string   `json:"liquidityPoolToken" bson:"liquidityPoolToken"`
	PTokensAddrs    []string `json:"provisionTokens" bson:"provisionTokens"`
}

type StakingConfig struct {
	Pools []*Pool `json:"pools" bson:"pools"`
}

type Action struct {
	Name          string         `json:"name" bson:"name"`
	LPTokenAmount AmountBigFloat `json:"lpTokenAmount" bson:"lpTokenAmount"`
}

type LPAssets struct {
	LPToken           *Token                    `json:"liquidityToken" bson:"liquidityToken"`
	ReserveTokens     []*Token                  `json:"reserveTokens" bson:"reserveTokens"`
	LPTokenToReserves map[string]AmountBigFloat `json:"lpTokenToReserves" bson:"lpTokenToReserves"`
}

type LPool struct {
	A string `json:"a" bson:"a"`
}

type TxData struct {
	// Hash   common.Hash    `json:"hash" bson:"hash"`
	// From   common.Address `json:"from" bson:"from"`
	// To     common.Address `json:"to" bson:"to"`
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

	Protocol      string    `json:"protocol" bson:"protocol"`
	LiquidityPool LPool     `json:"liquidityPool" bson:"liquidityPool"`
	Timestamp     time.Time `json:"timestamp" bson:"timestamp"`
	Action        *Action   `json:"action" bson:"action"`
	LPAssets      *LPAssets `json:"liquidityPoolAssets" bson:"liquidityPoolAssets"`
}
