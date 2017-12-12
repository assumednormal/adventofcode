package advent2017

import (
	"strconv"
	"strings"
)

// GroupWithZero http://adventofcode.com/2017/day/12
func GroupWithZero(s string) int {
	// parse input
	cnxs := make(map[int][]int)
	for _, l := range strings.Split(s, "\n") {
		t := strings.Split(l, " <-> ")
		src, err := strconv.Atoi(t[0])
		if err != nil {
			panic(err)
		}
		trgts := []int{}
		for _, target := range strings.Split(t[1], ", ") {
			v, err := strconv.Atoi(target)
			if err != nil {
				panic(err)
			}
			trgts = append(trgts, v)
		}
		cnxs[src] = trgts
	}

	// find connections to 0
	done := make(map[int]bool)
	done[0] = false
	for {
		// which src is next?
		nextSrc := -1
		for src, d := range done {
			if !d {
				nextSrc = src
				break
			}
		}
		if nextSrc == -1 {
			break
		}
		for _, v := range cnxs[nextSrc] {
			if !done[v] {
				done[v] = false
			}
		}
		done[nextSrc] = true
	}

	return len(done)
}

func findGroup(start int, cnxs map[int][]int) []int {
	// find connections to start
	done := make(map[int]bool)
	done[start] = false
	for {
		// which src is next?
		nextSrc := -1
		for src, d := range done {
			if !d {
				nextSrc = src
				break
			}
		}
		if nextSrc == -1 {
			break
		}
		for _, v := range cnxs[nextSrc] {
			if !done[v] {
				done[v] = false
			}
		}
		done[nextSrc] = true
	}

	res := []int{}
	for k := range done {
		res = append(res, k)
	}

	return res
}

// CountGroups http://adventofcode.com/2017/day/12
func CountGroups(s string) int {
	// parse input
	cnxs := make(map[int][]int)
	for _, l := range strings.Split(s, "\n") {
		t := strings.Split(l, " <-> ")
		src, err := strconv.Atoi(t[0])
		if err != nil {
			panic(err)
		}
		trgts := []int{}
		for _, target := range strings.Split(t[1], ", ") {
			v, err := strconv.Atoi(target)
			if err != nil {
				panic(err)
			}
			trgts = append(trgts, v)
		}
		cnxs[src] = trgts
	}

	groups := 0
	seen := []int{}
	// for each src, determine its group
	for k := range cnxs {
		inSeen := false
		for _, v := range seen {
			if k == v {
				inSeen = true
			}
		}
		if !inSeen {
			groups++
			seen = append(seen, findGroup(k, cnxs)...)
		}
	}

	return groups
}
