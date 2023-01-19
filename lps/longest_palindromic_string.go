package lps

import (
	"fmt"
	"strings"
)

func LongestPalindrome(s string) string {
	if s[0] == s[len(s)-1] {
		if reverse(s) == s {
			return s
		}
	}
	ss := strings.Split(s, "")
	ans := ss[0]
	for i := 0; i < len(ss); i++ {

		f := ss[i]
		pans := ss[i]

		for j := i + 1; j < len(ss); j++ {
			pans = fmt.Sprintf("%s%s", pans, ss[j])

			if ss[j] == f && reverse(pans) == pans && len(pans) > len(ans) {
				ans = pans
			}
		}

	}

	return ans
}

func reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}
