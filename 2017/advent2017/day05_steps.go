package advent2017

import (
	"strconv"
	"strings"
)

// StepsToExit http://adventofcode.com/2017/day/5
func StepsToExit(s string) int {
	// convert string to []int
	p := []int{}
	for _, u := range strings.Fields(s) {
		v, err := strconv.Atoi(u)
		if err != nil {
			panic(err)
		}
		p = append(p, v)
	}

	// follow path
	pos := 0
	steps := 0
	for {
		next := p[pos]
		p[pos]++
		steps++
		if len(p) <= pos+next {
			break
		}
		pos += next
	}

	return steps
}

// StepsToExit2 http://adventofcode.com/2017/day/5
func StepsToExit2(s string) int {
	// convert string to []int
	p := []int{}
	for _, u := range strings.Fields(s) {
		v, err := strconv.Atoi(u)
		if err != nil {
			panic(err)
		}
		p = append(p, v)
	}

	// follow path
	pos := 0
	steps := 0
	for {
		next := p[pos]
		if next >= 3 {
			p[pos]--
		} else {
			p[pos]++
		}
		steps++
		if len(p) <= pos+next {
			break
		}
		pos += next
	}

	return steps
}
