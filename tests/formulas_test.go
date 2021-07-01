package tests

import (
	"blockchain-data-collector/formulas"
	"fmt"
	"math/big"
	"testing"
)

func TestCalculateImpermanentLoss(t *testing.T) {
	total := formulas.CalculateImpermanentLoss(new(big.Float).SetFloat64(1))

	fmt.Print(total)

	if total.Cmp(new(big.Float).SetFloat64(0)) != 0 {
		t.Errorf("Sum was incorrect, got: %d, want: %d.", total, 0)
	}
}

func TestCalculateImpermanentLossPercentage(t *testing.T) {
	total := formulas.CalculateImpermanentLossPercentage(new(big.Float).SetFloat64(1))

	fmt.Print(total)

	if total.Cmp(new(big.Float).SetFloat64(100)) != 0 {
		t.Errorf("Sum was incorrect, got: %d, want: %d.", total, 100)
	}
}

func TestCalculateImpermanentLossUSD(t *testing.T) {
	total := formulas.CalculateImpermanentLossUSD(new(big.Float).SetFloat64(2), new(big.Float).SetFloat64(2))

	fmt.Print(total)

	if total.Cmp(new(big.Float).SetFloat64(4)) != 0 {
		t.Errorf("Sum was incorrect, got: %d, want: %d.", total, 4)
	}
}

func TestCalculateNetGainLoss(t *testing.T) {
	total := formulas.CalculateNetGainLoss(new(big.Float).SetFloat64(2), new(big.Float).SetFloat64(2))

	fmt.Print(total)

	if total.Cmp(new(big.Float).SetFloat64(0)) != 0 {
		t.Errorf("Sum was incorrect, got: %d, want: %d.", total, 0)
	}
}

func TestCalculateAssetPrice(t *testing.T) {
	total := formulas.CalculateAssetPrice(new(big.Float).SetFloat64(2), new(big.Float).SetFloat64(2))

	fmt.Print(total)

	if total.Cmp(new(big.Float).SetFloat64(4)) != 0 {
		t.Errorf("Sum was incorrect, got: %d, want: %d.", total, 4)
	}
}

func TestCalculateAssetsValue(t *testing.T) {
	total := formulas.CalculateAssetsValue([]*big.Float{new(big.Float).SetFloat64(2), new(big.Float).SetFloat64(2)})

	fmt.Print(total)

	if total.Cmp(new(big.Float).SetFloat64(4)) != 0 {
		t.Errorf("Sum was incorrect, got: %d, want: %d.", total, 4)
	}
}

func TestCalculateFees(t *testing.T) {
	total := formulas.CalculateFees(new(big.Float).SetFloat64(3), new(big.Float).SetFloat64(2), new(big.Float).SetFloat64(1))

	fmt.Print(total)

	if total.Cmp(new(big.Float).SetFloat64(0)) != 0 {
		t.Errorf("Sum was incorrect, got: %d, want: %d.", total, 0)
	}
}
