package advent2017

import (
	"strings"
)

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// HexDistance http://adventofcode.com/2017/day/11
func HexDistance(s string) int {
	moves := strings.Split(s, ",")

	aggMoves := make(map[string]int)
	for _, m := range moves {
		aggMoves[m]++
	}

	// ne and sw cancel
	neswMin := min(aggMoves["ne"], aggMoves["sw"])
	aggMoves["ne"] -= neswMin
	aggMoves["sw"] -= neswMin

	// nw and se cancel
	nwseMin := min(aggMoves["nw"], aggMoves["se"])
	aggMoves["nw"] -= nwseMin
	aggMoves["se"] -= nwseMin

	// ne and nw combine to n
	nenwMin := min(aggMoves["ne"], aggMoves["nw"])
	aggMoves["ne"] -= nenwMin
	aggMoves["nw"] -= nenwMin
	aggMoves["n"] += nenwMin

	// se and sw combine to s
	seswMin := min(aggMoves["se"], aggMoves["sw"])
	aggMoves["se"] -= seswMin
	aggMoves["sw"] -= seswMin
	aggMoves["s"] += seswMin

	// n and s cancel
	nsMin := min(aggMoves["n"], aggMoves["s"])
	aggMoves["n"] -= nsMin
	aggMoves["s"] -= nsMin

	res := 0
	if aggMoves["n"] > 0 && (aggMoves["ne"] > 0 || aggMoves["nw"] > 0) {
		res = aggMoves["n"] + aggMoves["ne"] + aggMoves["nw"]
	} else if aggMoves["s"] > 0 && (aggMoves["se"] > 0 || aggMoves["sw"] > 0) {
		res = aggMoves["s"] + aggMoves["se"] + aggMoves["sw"]
	} else if aggMoves["n"] > 0 && (aggMoves["se"] > 0 || aggMoves["sw"] > 0) {
		res = max(aggMoves["n"], aggMoves["se"]+aggMoves["sw"])
	} else if aggMoves["s"] > 0 && (aggMoves["ne"] > 0 || aggMoves["nw"] > 0) {
		res = max(aggMoves["s"], aggMoves["ne"]+aggMoves["nw"])
	} else {
		res = aggMoves["ne"] + aggMoves["se"] + aggMoves["sw"] + aggMoves["se"]
	}

	return res
}

func copyMap(m map[string]int) map[string]int {
	n := make(map[string]int)
	for k, v := range m {
		n[k] = v
	}
	return n
}

func distance(m map[string]int) int {
	// ne and sw cancel
	neswMin := min(m["ne"], m["sw"])
	m["ne"] -= neswMin
	m["sw"] -= neswMin

	// nw and se cancel
	nwseMin := min(m["nw"], m["se"])
	m["nw"] -= nwseMin
	m["se"] -= nwseMin

	// ne and nw combine to n
	nenwMin := min(m["ne"], m["nw"])
	m["ne"] -= nenwMin
	m["nw"] -= nenwMin
	m["n"] += nenwMin

	// se and sw combine to s
	seswMin := min(m["se"], m["sw"])
	m["se"] -= seswMin
	m["sw"] -= seswMin
	m["s"] += seswMin

	// n and s cancel
	nsMin := min(m["n"], m["s"])
	m["n"] -= nsMin
	m["s"] -= nsMin

	res := 0
	if m["n"] > 0 && (m["ne"] > 0 || m["nw"] > 0) {
		res = m["n"] + m["ne"] + m["nw"]
	} else if m["s"] > 0 && (m["se"] > 0 || m["sw"] > 0) {
		res = m["s"] + m["se"] + m["sw"]
	} else if m["n"] > 0 && (m["se"] > 0 || m["sw"] > 0) {
		res = max(m["n"], m["se"]+m["sw"])
	} else if m["s"] > 0 && (m["ne"] > 0 || m["nw"] > 0) {
		res = max(m["s"], m["ne"]+m["nw"])
	} else {
		res = m["ne"] + m["se"] + m["sw"] + m["se"]
	}

	return res
}

// HexDistance2 http://adventofcode.com/2017/day/11
func HexDistance2(s string) int {
	moves := strings.Split(s, ",")

	maxDist := 0

	aggMoves := make(map[string]int)
	for _, m := range moves {
		aggMoves[m]++
		maxDist = max(maxDist, distance(copyMap(aggMoves)))
	}

	return maxDist
}
