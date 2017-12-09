package advent2017

import (
	"sort"
	"strings"
)

// ValidPassphrases http://adventofcode.com/2017/day/4
func ValidPassphrases(s string) int {
	t := 0
	c := 0
	for _, r := range strings.Split(s, "\n") {
		t++
		m := make(map[string]bool)
		for _, p := range strings.Fields(r) {
			if _, ok := m[p]; ok {
				c++
				break
			}
			m[p] = true
		}
	}
	return t - c
}

// ValidAnagramPassphrases http://adventofcode.com/2017/day/4
func ValidAnagramPassphrases(s string) int {
	t := 0
	c := 0
	for _, r := range strings.Split(s, "\n") {
		t++
		m := make(map[string]bool)
		for _, p := range strings.Fields(r) {
			q := strings.Split(p, "")
			sort.Strings(q)
			p = strings.Join(q, "")
			if _, ok := m[p]; ok {
				c++
				break
			}
			m[p] = true
		}
	}
	return t - c
}
