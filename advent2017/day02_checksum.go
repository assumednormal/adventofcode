package advent2017

import (
	"math"
	"strconv"
	"strings"
)

// RowMaxDiffChecksum http://adventofcode.com/2017/day/2
func RowMaxDiffChecksum(s string) int {
	var checksum int
	for _, r := range strings.Split(s, "\n") {
		low := math.MaxInt64
		high := math.MinInt64
		for _, c := range strings.Fields(r) {
			v, err := strconv.Atoi(c)
			if err != nil {
				panic(err)
			}
			if v < low {
				low = v
			}
			if high < v {
				high = v
			}
		}
		checksum += high - low
	}
	return checksum
}
