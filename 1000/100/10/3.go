package leetcode_10

import "leetcode/util"

// Time complexity: O(n^3)
// Space complexity: O(n)
func lengthOfLongestSubstring_bruteForce(s string) int {
	n := len(s)
	res := 0

	for i := 0; i < n; i++ {
		for j := i; j < n; j++ {
			if checkDuplicateChar(i, j, s) {
				res = util.Max(res, j-i+1)
			} else {
				break
			}
		}
	}

	return res
}

// Helper function to check if all characters in the substring s[i:j+1] are unique
func checkDuplicateChar(start, end int, s string) bool {
	m := make(map[byte]bool)

	for i := start; i <= end; i++ {
		c := s[i]
		if m[c] {
			return false
		}
		m[c] = true
	}
	return true
}

// Time complexity: O(n)
// Space complexity: O(n)
func lengthOfLongestSubstring_slidingWindow(s string) int {
	charSet := make(map[byte]bool)
	l := 0
	res := 0

	for r := range s {
		for charSet[s[r]] {
			delete(charSet, s[l])
			l++
		}
		charSet[s[r]] = true
		res = util.Max(res, r-l+1)
	}
	return res
}
