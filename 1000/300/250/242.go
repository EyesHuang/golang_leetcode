package leetcode_250

import (
	"sort"
	"strings"
)

// Time complexity: O(n^2)
// Space complexity: O(n)
func isAnagram_bruteForce(s, t string) bool {
	sArr := []rune(s)
	tArr := []rune(t)

	if len(sArr) != len(tArr) {
		return false
	}

	for _, sEle := range sArr {
		for tIdx, tEle := range tArr {
			if sEle == tEle {
				tArr = append(tArr[:tIdx], tArr[(tIdx+1):]...)
				break
			}
		}
	}

	if len(tArr) == 0 {
		return true
	} else {
		return false
	}
}

// Time complexity: O(n)
// Space complexity: O(1)
func isAnagram_frequencyCounter(s, t string) bool {
	if len(s) != len(t) {
		return false
	}

	table := make([]int, 26)

	for i := 0; i < len(s); i++ {
		table[s[i]-'a']++
	}

	for i := 0; i < len(t); i++ {
		table[t[i]-'a']--

		if table[t[i]-'a'] < 0 {
			return false
		}
	}

	return true
}

// Time complexity: O(nlogn)
// Space complexity: O(n)
func isAnagram_sorting(s, t string) bool {
	if len(s) != len(t) {
		return false
	}

	runeS := []rune(s)
	runeT := []rune(t)

	sort.Slice(runeS, func(i, j int) bool { return runeS[i] < runeS[j] })
	sort.Slice(runeT, func(i, j int) bool { return runeT[i] < runeT[j] })

	return strings.Compare(string(runeS), string(runeT)) == 0
}
