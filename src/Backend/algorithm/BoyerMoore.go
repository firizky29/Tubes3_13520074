package algorithm

func badCharHeuristic(str string, size int, badchar []int) {
	var i int
	for i = 0; i < size; i++ {
		badchar[i] = -1
	}
	for i = 0; i < size; i++ {
		badchar[int(str[i])] = i
	}
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func search(human string, virus string, flag *string) {
	*flag = "not found"
	var m = len(virus)
	var n = len(human)
	var s = 0
	var badchar = make([]int, 256)

	badCharHeuristic(virus, m, badchar)

	for s <= (n - m) {
		var j = m - 1

		for j >= 0 && virus[j] == human[s+j] {
			j--
		}
		if j < 0 {
			*flag = "found"
			if s+m < n {
				s += m - badchar[human[s+m]]
			} else {
				s += 1
			}
		} else {
			s += max(1, j-badchar[human[s+j]])
		}
	}
}
