package valid_palindrome

import (
	"strings"
	"unicode"
)

// Time complexity: O(n)
// Space complexity: O(n)
func isPalindrome_reverse(s string) bool {
	var filtered string

	for _, char := range s {
		if unicode.IsLetter(char) || unicode.IsNumber(char) {
			filtered += strings.ToLower(string(char))
		}
	}

	chars := []rune(filtered)

	for i, j := 0, len(filtered)-1; i < j; i, j = i+1, j-1 {
		chars[i], chars[j] = chars[j], chars[i]
	}

	return filtered == string(chars)
}

// Time complexity: O(n)
// Space complexity: O(1)
func isPalindrome_twoPointers(s string) bool {
	f := func(r rune) rune {
		if !unicode.IsLetter(r) && !unicode.IsNumber(r) {
			return -1
		}

		return unicode.ToLower(r)
	}

	s = strings.Map(f, s)

	i, j := 0, len(s)-1

	for i < j {
		if s[i] != s[j] {
			return false
		}

		i = i + 1
		j = j - 1
	}

	return true
}
