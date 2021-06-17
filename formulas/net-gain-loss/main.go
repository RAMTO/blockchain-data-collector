package netgainloss

import "math/big"

func CalculateNetGainLoss(initialAssetsInitialPrices, currentAssestsCurrentPrices *big.Float) *big.Float {
	finalRestult := new(big.Float).Sub(currentAssestsCurrentPrices, initialAssetsInitialPrices)

	return finalRestult
}
