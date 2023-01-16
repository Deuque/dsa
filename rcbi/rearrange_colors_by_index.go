package rcbi

import (
	"strconv"
	"strings"
)

// [green4, red1, blue5] should give [red, green, blue]
func Solution(s string) string {
	colors := strings.Split(s, " ")
	scolors := []string{}

	for i := 0; i < len(colors); i++ {
		scolors = append(scolors, "")
	}

	for _, c := range colors {
		litem := c[len(c)-1:]
		litemInt, _ := strconv.Atoi(litem)
		scolors[litemInt-1] = c[0 : len(c)-1]
	}

	return strings.Join(scolors, " ")

}
