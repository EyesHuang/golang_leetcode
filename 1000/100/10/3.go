package leetcode_10

func LengthOfLongestSubstringBruteForce(s string) int {
	n := len(s)
	res := 0

	for i := 0; i < n; i++ {
		for j := i; j < n; j++ {
			if checkDuplicateChar(i, j, s) {
				res = max(res, j-i+1)
			} else {
				break
			}
		}
	}

	return res
}

func checkDuplicateChar(start, end int, s string) bool {
	m := make(map[uint8]int)

	for i := start; i <= end; i++ {
		c := s[i]
		if _, isExist := m[c]; isExist {
			return false
		}
		m[c]++
	}
	return true
}

func LengthOfLongestSubstringS1(s string) int {
	counts := make(map[uint8]int)
	left := 0
	right := 0
	res := 0

	for i := 0; i < len(s); i++ {
		c := s[i]

		if value, isExist := counts[c]; isExist {
			left = max(left, value+1)
		}

		counts[c] = right
		res = max(res, right-left+1)

		right++
	}

	return res
}

func LengthOfLongestSubstringS2(s string) int {
	m := make(map[uint8]int)
	i, maxLen := 0, 0

	for i < len(s) {
		c := s[i]
		if val, isExist := m[c]; !isExist {
			m[c] = i
			i++
		} else {
			maxLen = max(maxLen, len(m))
			i = val + 1
			m = make(map[uint8]int)
		}
	}

	maxLen = max(maxLen, len(m))
	m = nil
	return maxLen
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
