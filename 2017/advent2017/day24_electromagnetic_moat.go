package advent2017

import (
	"strconv"
	"strings"
)

type component struct {
	port1, port2 int
}

func hasNPins(c component, n int) bool {
	return c.port1 == n || c.port2 == n
}

func nFirst(c component, n int) component {
	if c.port1 == n {
		return component{c.port1, c.port2}
	}
	return component{c.port2, c.port1}
}

// Bridge https://adventofcode.com/2017/day/24
func Bridge(s string) int {
	components := []component{}
	for _, l := range strings.Split(s, "\n") {
		vals := strings.Split(l, "/")
		a, err := strconv.Atoi(vals[0])
		if err != nil {
			panic(err)
		}
		b, err := strconv.Atoi(vals[1])
		if err != nil {
			panic(err)
		}
		components = append(components, component{a, b})
	}

	return maxBridgeStrength(component{0, 0}, components)
}

func maxBridgeStrength(last component, remaining []component) int {
	var max int
	for i, c := range remaining {
		if hasNPins(c, last.port2) {
			cpy := make([]component, len(remaining)-1)
			copy(cpy, remaining[:i])
			copy(cpy[i:], remaining[i+1:])
			tmp := maxBridgeStrength(nFirst(c, last.port2), cpy)
			if max < tmp {
				max = tmp
			}
		}
	}
	return last.port1 + last.port2 + max
}

// Bridge2 https://adventofcode.com/2017/day/24
func Bridge2(s string) int {
	components := []component{}
	for _, l := range strings.Split(s, "\n") {
		vals := strings.Split(l, "/")
		a, err := strconv.Atoi(vals[0])
		if err != nil {
			panic(err)
		}
		b, err := strconv.Atoi(vals[1])
		if err != nil {
			panic(err)
		}
		components = append(components, component{a, b})
	}

	c, _ := maxBridgeLengthStrength(component{0, 0}, components)
	return c
}

func maxBridgeLengthStrength(last component, remaining []component) (int, int) {
	var maxStrength, maxLength int
	for i, c := range remaining {
		if hasNPins(c, last.port2) {
			cpy := make([]component, len(remaining)-1)
			copy(cpy, remaining[:i])
			copy(cpy[i:], remaining[i+1:])
			tmpStrength, tmpLength := maxBridgeLengthStrength(nFirst(c, last.port2), cpy)
			if maxLength == tmpLength && maxStrength < tmpStrength {
				maxStrength = tmpStrength
				maxLength = tmpLength
			} else if maxLength < tmpLength {
				maxStrength = tmpStrength
				maxLength = tmpLength
			}
		}
	}
	return last.port1 + last.port2 + maxStrength, 1 + maxLength
}
