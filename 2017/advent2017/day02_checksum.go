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

// RowDivChecksum http://adventofcode.com/2017/day/2
func RowDivChecksum(s string) int {
	m := make([][]int, 0)
	for _, r := range strings.Split(s, "\n") {
		u := make([]int, 0)
		for _, c := range strings.Fields(r) {
			v, err := strconv.Atoi(c)
			if err != nil {
				panic(err)
			}
			u = append(u, v)
		}
		m = append(m, u)
	}

	var checksum int
	for _, r := range m {
		for j := 0; j < len(r)-1; j++ {
			for k := j + 1; k < len(r); k++ {
				if r[k]%r[j] == 0 {
					checksum += r[k] / r[j]
				} else if r[j]%r[k] == 0 {
					checksum += r[j] / r[k]
				}
			}
		}
	}
	return checksum
}
