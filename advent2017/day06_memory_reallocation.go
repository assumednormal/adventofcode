package advent2017

import (
	"fmt"
	"strconv"
	"strings"
)

// StepsToPriorPattern http://adventofcode.com/2017/day/6
func StepsToPriorPattern(s string) int {
	// convert string to []int
	p := []int{}
	for _, u := range strings.Fields(s) {
		v, err := strconv.Atoi(u)
		if err != nil {
			panic(err)
		}
		p = append(p, v)
	}

	// memory reallocations
	m := make(map[string]bool)
	m[fmt.Sprint(p)] = true
	steps := 0

	for {
		steps++

		// find position with max blocks
		whichMax := 0
		for i, v := range p {
			if v > p[whichMax] {
				whichMax = i
			}
		}

		// reallocate
		blocks := p[whichMax]
		p[whichMax] = 0
		for i := 0; i < blocks; i++ {
			p[(whichMax+i+1)%len(p)]++
		}

		// check for matching pattern
		if _, ok := m[fmt.Sprint(p)]; ok {
			break
		}
		m[fmt.Sprint(p)] = true
	}

	return steps
}

// StepsToPriorPattern2 http://adventofcode.com/2017/day/6
func StepsToPriorPattern2(s string) int {
	// convert string to []int
	p := []int{}
	for _, u := range strings.Fields(s) {
		v, err := strconv.Atoi(u)
		if err != nil {
			panic(err)
		}
		p = append(p, v)
	}

	// memory reallocations
	m := make(map[string]int)
	steps := 0
	m[fmt.Sprint(p)] = steps
	diff := 0

	for {
		steps++

		// find position with max blocks
		whichMax := 0
		for i, v := range p {
			if v > p[whichMax] {
				whichMax = i
			}
		}

		// reallocate
		blocks := p[whichMax]
		p[whichMax] = 0
		for i := 0; i < blocks; i++ {
			p[(whichMax+i+1)%len(p)]++
		}

		// check for matching pattern
		if seen, ok := m[fmt.Sprint(p)]; ok {
			diff = steps - seen
			break
		}
		m[fmt.Sprint(p)] = steps
	}

	return diff
}
