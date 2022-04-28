package algorithm

import (
	"github.com/shopspring/decimal"
)

func lcs(virus, human string) int {
	virusLen := len(virus)
	humanLen := len(human)
	result := 0
	var suffix = make([][]int, virusLen+1)
	for i := range suffix {
		suffix[i] = make([]int, humanLen+1)
	}

	for i := 0; i <= virusLen; i++ {
		for j := 0; j <= humanLen; j++ {
			if i == 0 || j == 0 {
				suffix[i][j] = 0
			} else if virus[i-1] == human[j-1] {
				suffix[i][j] = 1 + suffix[i-1][j-1]
				if suffix[i][j] > result {
					result = suffix[i][j]
				}
			} else {
				suffix[i][j] = max(suffix[i-1][j], suffix[i][j-1])
			}
		}
	}
	return result
}

func SimilarityLevel(virus, human string) decimal.Decimal {
	return decimal.NewFromFloat(float64((lcs(virus, human) * 100.0) / len(virus)))
}
