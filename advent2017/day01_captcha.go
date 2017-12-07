package advent2017

import (
	"fmt"
	"strconv"
)

// SumConsecutiveDigits from http://adventofcode.com/2017/day/1
func SumConsecutiveDigits(s string) int {
	t := 0
	for i := range []byte(s) {
		if s[i%len(s)] == s[(i+1)%len(s)] {
			u, err := strconv.Atoi(fmt.Sprintf("%c", s[i%len(s)]))
			if err != nil {
				panic(err)
			}
			t += u
		}
	}
	return t
}
