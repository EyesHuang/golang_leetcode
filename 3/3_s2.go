package longest_substring_without_repeating_characters

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func LengthOfLongestSubstring(s string) int {
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
