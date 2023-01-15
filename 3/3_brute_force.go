package longest_substring_without_repeating_characters

func maxBruteforce(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func LengthOfLongestSubstringBruteForce(s string) int {
	n := len(s)
	res := 0

	for i := 0; i < n; i++ {
		for j := i; j < n; j++ {
			if checkDuplicateChar(i, j, s) {
				res = maxBruteforce(res, j-i+1)
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
