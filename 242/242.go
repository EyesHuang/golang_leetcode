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

func isAnagram_FrequencyCounter(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	counts := make(map[rune]int)

	for _, char := range s {
		counts[char]++
	}

	for _, char := range t {
		counts[char]--

		if counts[char] < 0 {
			return false
		}
	}

	return true
}
