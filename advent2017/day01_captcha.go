package advent2017

import (
	"fmt"
	"strconv"
)

// SumSkipDigits from http://adventofcode.com/2017/day/1
func SumSkipDigits(s string, d int) int {
	t := 0
	for i := range []byte(s) {
		if s[i%len(s)] == s[(i+d)%len(s)] {
			u, err := strconv.Atoi(fmt.Sprintf("%c", s[i%len(s)]))
			if err != nil {
				panic(err)
			}
			t += u
		}
	}
	return t
}
