package lsswrc

import (
	"fmt"
	"strings"
)

func LengthOfLongestSubstring(s string) int {
	if len(s) == 0 {
		return 0
	}
	if len(s) == 1 {
		return 1
	}
	count := 0

	sub := ""
	ssplit := strings.Split(s, "")
	for i, r := range ssplit {
		if strings.Contains(sub, r) {
			sublength := len(sub)
			if sublength > count {
				count = sublength
			}
			ns := s[strings.Index(s, r)+1:]
			fcount := LengthOfLongestSubstring(ns)
			if fcount > count {
				return fcount
			}
			return count

		} else {
			sub = fmt.Sprintf("%s%s", sub, r)
			if (i + 1) == len(ssplit) {
				sublength := len(sub)
				if sublength > count {
					count = sublength
				}
			}
		}
	}
	return count
}
