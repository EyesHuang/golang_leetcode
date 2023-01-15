package longest_substring_without_repeating_characters

func maxS1(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func LengthOfLongestSubstringS1(s string) int {
	counts := make(map[uint8]int)
	left := 0
	right := 0
	res := 0

	for i := 0; i < len(s); i++ {
		c := s[i]

		if value, isExist := counts[c]; isExist {
			left = maxS1(left, value+1)
		}

		counts[c] = right
		res = maxS1(res, right-left+1)

		right++
	}

	return res
}
