package valid_anagram

func isAnagram_bruteForce(s string, t string) bool {
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

func isAnagram_frequencyCounter(s, t string) bool {
	if len(s) != len(t) {
		return false
	}

	table := make([]int, 26)

	for i := 0; i < len(s); i++ {
		table[s[i]-'a']++
	}

	for i := 0; i < len(s); i++ {
		table[t[i]-'a']--

		if table[t[i]-'a'] < 0 {
			return false
		}
	}

	return true
}
