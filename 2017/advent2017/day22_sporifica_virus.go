package advent2017

import "strings"

type vState struct {
	grid       map[location]int
	l          location
	vi         int
	vj         int
	infections int
}

func newVState(s string) *vState {
	// fill grid with
	grid := make(map[location]int)
	var rows, cols int
	for i, row := range strings.Split(s, "\n") {
		for j, col := range row {
			if col == '#' {
				grid[location{i, j}] = 2
			}

			if i == 0 {
				cols++
			}
		}

		rows++
	}

	start := location{(rows - 1) / 2, (cols - 1) / 2}

	return &vState{grid, start, -1, 0, 0}
}

func (v *vState) burst() {
	// determine next direction
	if v.grid[v.l] == 0 {
		// if clean, turn left
		if v.vi == -1 {
			// we were looking up, now go left
			v.vi = 0
			v.vj = -1
		} else if v.vi == 1 {
			// we were looking down, now go right
			v.vi = 0
			v.vj = 1
		} else if v.vj == -1 {
			// we were looking left, now go down
			v.vi = 1
			v.vj = 0
		} else if v.vj == 1 {
			// we were looking right, now go up
			v.vi = -1
			v.vj = 0
		}
	} else if v.grid[v.l] == 2 {
		// if infected, turn right
		if v.vi == -1 {
			// we were looking up, now go right
			v.vi = 0
			v.vj = 1
		} else if v.vi == 1 {
			// we were looking down, now go left
			v.vi = 0
			v.vj = -1
		} else if v.vj == -1 {
			// we were looking left, now go up
			v.vi = -1
			v.vj = 0
		} else if v.vj == 1 {
			// we were looking right, now go down
			v.vi = 1
			v.vj = 0
		}
	} else if v.grid[v.l] == 3 {
		// if flagged, turn around
		v.vi = -1 * v.vi
		v.vj = -1 * v.vj
	}

	// change state
	v.grid[v.l] = (v.grid[v.l] + 1) % 4
	if v.grid[v.l] == 2 {
		v.infections++
	}

	// move
	v.l = location{v.l.I + v.vi, v.l.J + v.vj}
}

// Infections http://adventofcode.com/2017/day/22
func Infections(s string, i int) int {
	v := newVState(s)
	for j := 0; j < i; j++ {
		v.burst()
	}
	return v.infections
}

// Infections2 http://adventofcode.com/2017/day/22
func Infections2(s string, i int) int {
	v := newVState(s)
	for j := 0; j < i; j++ {
		v.burst()
	}
	return v.infections
}
