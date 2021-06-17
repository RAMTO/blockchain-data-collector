package impermanentloss

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
