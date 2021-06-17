package tradingfees

import (
	"math/big"
)

func CalculateAssetPrice(asset, price *big.Float) *big.Float {
	totalPrice := new(big.Float).Mul(asset, price)

	return totalPrice
}

func CalculateAssetsValue(assets []*big.Float) *big.Float {
	result := big.NewFloat(float64(0))

	for _, asset := range assets {
		result = new(big.Float).Add(result, asset)
	}

	return result
}

func CalculateFees(assetsCurrentValue, assetsInitialValue, impermanentloss *big.Float) *big.Float {
	priceDiff := new(big.Float).Sub(assetsCurrentValue, assetsInitialValue)
	finalRestult := new(big.Float).Sub(priceDiff, impermanentloss)

	return finalRestult
}
