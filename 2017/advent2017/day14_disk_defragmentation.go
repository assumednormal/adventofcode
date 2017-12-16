package advent2017

import (
	"fmt"
)

func hexToBin(s string) string {
	m := make(map[rune]string)
	m['0'] = "0000"
	m['1'] = "0001"
	m['2'] = "0010"
	m['3'] = "0011"
	m['4'] = "0100"
	m['5'] = "0101"
	m['6'] = "0110"
	m['7'] = "0111"
	m['8'] = "1000"
	m['9'] = "1001"
	m['a'] = "1010"
	m['b'] = "1011"
	m['c'] = "1100"
	m['d'] = "1101"
	m['e'] = "1110"
	m['f'] = "1111"

	t := ""
	for _, r := range s {
		t = fmt.Sprintf("%s%s", t, m[r])
	}
	return t
}

// UsedSquares http://adventofcode.com/2017/day/14
func UsedSquares(s string) int {
	used := 0
	for i := 0; i < 128; i++ {
		kn := KnotHash2([]byte(fmt.Sprintf("%s-%d", s, i)))
		for _, r := range hexToBin(kn) {
			if r == '1' {
				used++
			}
		}
	}
	return used
}

type location struct {
	I, J int
}

func neighbors(a location, m map[location]bool) map[location]bool {
	frontier := make(map[location]bool)
	frontier[a] = true
	seen := make(map[location]bool)
	for len(frontier) > 0 {
		var next location
		for k := range frontier {
			next = k
			break
		}
		delete(frontier, next)
		seen[next] = true

		if m[location{next.I, next.J - 1}] && !seen[location{next.I, next.J - 1}] {
			frontier[location{next.I, next.J - 1}] = true
		}
		if m[location{next.I, next.J + 1}] && !seen[location{next.I, next.J + 1}] {
			frontier[location{next.I, next.J + 1}] = true
		}
		if m[location{next.I - 1, next.J}] && !seen[location{next.I - 1, next.J}] {
			frontier[location{next.I - 1, next.J}] = true
		}
		if m[location{next.I + 1, next.J}] && !seen[location{next.I + 1, next.J}] {
			frontier[location{next.I + 1, next.J}] = true
		}
	}
	return seen
}

// CountRegions http://adventofcode.com/2017/day/14
func CountRegions(s string) int {
	grid := make(map[location]bool)
	for i := 0; i < 128; i++ {
		kn := KnotHash2([]byte(fmt.Sprintf("%s-%d", s, i)))
		for j, r := range hexToBin(kn) {
			grid[location{i, j}] = r == '1'
		}
	}

	regions := 0
	frontier := make(map[location]bool)
	for k, v := range grid {
		if v {
			frontier[k] = true
		}
	}
	seen := make(map[location]bool)
	for len(frontier) > 0 {
		var next location
		for k := range frontier {
			next = k
			break
		}
		regions++
		for l := range neighbors(next, frontier) {
			delete(frontier, l)
			seen[l] = true
		}
	}

	return regions
}
