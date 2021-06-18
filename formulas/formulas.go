package formulas

import "math/big"

func CalculateImpermanentLoss(r *big.Float) *big.Float {
	// 2*(math.Sqrt(r)/(1+r)) - 1
	oneBF := big.NewFloat(float64(1))
	twoBF := big.NewFloat(float64(2))
	rSqrt := new(big.Float).Sqrt(r)
	onePlusR := new(big.Float).Add(r, oneBF)
	rSqrtQuoOnePlusR := new(big.Float).Quo(rSqrt, onePlusR)
	twoMulRSqrtQuoOnePlusR := new(big.Float).Mul(twoBF, rSqrtQuoOnePlusR)
	result := new(big.Float).Sub(twoMulRSqrtQuoOnePlusR, oneBF)

	return result
}

func CalculateImpermanentLossPercentage(il *big.Float) *big.Float {
	hundredBF := big.NewFloat(float64(100))
	result := new(big.Float).Mul(il, hundredBF)

	return result
}

func CalculateImpermanentLossUSD(assets, il *big.Float) *big.Float {
	result := new(big.Float).Mul(assets, il)

	return result
}

func CalculateNetGainLoss(initialAssetsInitialPrices, currentAssestsCurrentPrices *big.Float) *big.Float {
	finalRestult := new(big.Float).Sub(currentAssestsCurrentPrices, initialAssetsInitialPrices)

	return finalRestult
}

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
