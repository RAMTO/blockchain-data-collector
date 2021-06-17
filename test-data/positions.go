package testdata

import (
	"blockchain-data-collector/models"
	"math/big"
)

func GetUserPositions() []*models.UserData {
	// Initial position
	UserPositions := []*models.UserData{
		{
			AssetAAmount:   new(big.Float).SetFloat64(68.10297729),
			AssetBAmount:   new(big.Float).SetFloat64(36.88957616),
			AssetAPriceUSD: new(big.Float).SetFloat64(0.404),
			AssetBPriceUSD: new(big.Float).SetFloat64(0.600),
		},
		{
			AssetAAmount:   new(big.Float).SetFloat64(141.6950545),
			AssetBAmount:   new(big.Float).SetFloat64(70.85011997),
			AssetAPriceUSD: new(big.Float).SetFloat64(0.41208),
			AssetBPriceUSD: new(big.Float).SetFloat64(0.594),
		},
		{
			AssetAAmount:   new(big.Float).SetFloat64(45.9772113),
			AssetBAmount:   new(big.Float).SetFloat64(20.40483455),
			AssetAPriceUSD: new(big.Float).SetFloat64(0.4444),
			AssetBPriceUSD: new(big.Float).SetFloat64(0.612),
		},
	}

	return UserPositions
}
