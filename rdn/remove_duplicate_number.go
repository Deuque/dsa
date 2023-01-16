package rdn

import (
	"fmt"

	"golang.org/x/exp/slices"
)

// [0, 0, 1, 1, 1, 2, 2, 3, 3, 4}] should give [0, 1, 2, 3, 4, _, _, _, _, _]
func Solution(arr []int) []string {
	foundDigits := make(map[int]bool)

	sortedArr := []string{}

	for _, d := range arr {
		sortedArr = append(sortedArr, "_")
		_, ok := foundDigits[d]
		if !ok {
			index := slices.Index(sortedArr, "_")
			sortedArr[index] = fmt.Sprintf("%v", d)
			foundDigits[d] = true
		}

	}

	return sortedArr
}
