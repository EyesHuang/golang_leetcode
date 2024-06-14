package leetcode_50

// Time complexity: O(KN)
// Space complexity: O(KN)
func groupAnagrams(strs []string) [][]string {
	type count [26]int
	m := make(map[count][]string)

	for _, str := range strs {
		var counts count
		for _, char := range str {
			counts[rune(char)-'a']++
		}

		m[counts] = append(m[counts], str)
	}

	res := make([][]string, 0, len(strs))
	for _, v := range m {
		res = append(res, v)
	}

	return res
}
