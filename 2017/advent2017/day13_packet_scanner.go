package advent2017

import (
	"strconv"
	"strings"
)

// TripSeverity http://adventofcode.com/2017/day/13
func TripSeverity(s string) int {
	// parse depths/ranges
	ranges := make(map[int]int)
	for _, l := range strings.Split(s, "\n") {
		t := strings.Split(l, ": ")
		d, err := strconv.Atoi(t[0])
		if err != nil {
			panic(err)
		}
		r, err := strconv.Atoi(t[1])
		if err != nil {
			panic(err)
		}
		ranges[d] = r
	}

	// find max trip length
	tripLength := 0
	for k := range ranges {
		tripLength = max(tripLength, k)
	}

	severity := 0
	for step := 0; step <= tripLength; step++ {
		// caught?
		if r, ok := ranges[step]; ok {
			if step%(2*r-2) == 0 {
				severity += step * r
			}
		}
	}

	return severity
}

func caught(delay, tripLength int, ranges map[int]int) bool {
	for step := 0; step <= tripLength; step++ {
		// caught?
		if r, ok := ranges[step]; ok {
			if (step+delay)%(2*r-2) == 0 {
				return true
			}
		}
	}

	return false
}

// MinimumDelay http://adventofcode.com/2017/day/13
func MinimumDelay(s string) int {
	// parse depths/ranges
	ranges := make(map[int]int)
	for _, l := range strings.Split(s, "\n") {
		t := strings.Split(l, ": ")
		d, err := strconv.Atoi(t[0])
		if err != nil {
			panic(err)
		}
		r, err := strconv.Atoi(t[1])
		if err != nil {
			panic(err)
		}
		ranges[d] = r
	}

	// find max trip length
	tripLength := 0
	for k := range ranges {
		tripLength = max(tripLength, k)
	}

	// search for first delay that has no severity
	delay := 0
	for {
		if !caught(delay, tripLength, ranges) {
			break
		}
		delay++
	}

	return delay
}
