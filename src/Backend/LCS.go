package main

func lcs(virus, human string, virLen, humLen int) int {
	if virLen == 0 || humLen == 0 {
		return 0
	}
	if virus[virLen-1] == human[humLen-1] {
		return 1 + lcs(virus, human, virLen-1, humLen-1)
	} else {
		return max(lcs(virus, human, virLen-1, humLen), lcs(virus, human, virLen, humLen-1))
	}
}
