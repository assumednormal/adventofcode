package advent2017

import (
	"errors"
	"fmt"
	"strings"
)

type state struct {
	grid  map[location]rune // grid
	rows  int               // number of rows in grid
	cols  int               // number of cols in grid
	seen  string            // letters we've seen
	l     location          // location in grio
	vi    int               // direction of travel in i direction
	vj    int               // direction of travel in j direction
	done  bool              // found end of path
	steps int               // number of steps taken
}

func newState(s string) (*state, error) {
	// parse grid
	grid := make(map[location]rune)
	rows := 0
	cols := 0
	for i, l := range strings.Split(s, "\n") {
		for j, r := range l {
			if i == 0 {
				cols++
			}
			if r == ' ' {
				continue
			}
			grid[location{i, j}] = r
		}
		rows++
	}

	// find starting point
	for j := 0; j < cols; j++ {
		if grid[location{0, j}] == '|' {
			return &state{grid, rows, cols, "", location{0, j}, 1, 0, false, 1}, nil
		}
	}

	return nil, errors.New("no starting point found")
}

func (s *state) String() string {
	return fmt.Sprintf("Current location: [%v]\nLetters Seen: [%v]\nHeading: [%d,%d]\n", s.l, s.seen, s.vi, s.vj)
}

func (s *state) solve() {
	// move in current direction until we hit a letter, a T, or blank (signaling that we're done)
	for {
		s.l = location{s.l.I + s.vi, s.l.J + s.vj}
		s.steps++
		switch {
		case 'A' <= s.grid[s.l] && s.grid[s.l] <= 'Z':
			s.seen += string(s.grid[s.l])
		case s.grid[s.l] == 0 || s.grid[s.l] == ' ':
			s.done = true
		case s.grid[s.l] == '+':
			// which direction should we go next?
			if s.vi != 0 {
				// we we re moving vertically, now go left or right
				l := location{s.l.I, s.l.J - 1}
				if ('A' <= s.grid[l] && s.grid[l] <= 'Z') || s.grid[l] == '-' || s.grid[l] == '+' {
					// go left
					s.vi = 0
					s.vj = -1
					break
				}
				l = location{s.l.I, s.l.J + 1}
				if ('A' <= s.grid[l] && s.grid[l] <= 'Z') || s.grid[l] == '-' || s.grid[l] == '+' {
					// go right
					s.vi = 0
					s.vj = 1
					break
				}
				// we can't go anywhere -> panic
				panic(fmt.Sprintf("I can't move from location [%v]", s.l))
			} else if s.vj != 0 {
				// we we re moving horizontally, now go up or down
				l := location{s.l.I - 1, s.l.J}
				if ('A' <= s.grid[l] && s.grid[l] <= 'Z') || s.grid[l] == '|' || s.grid[l] == '+' {
					// go up
					s.vi = -1
					s.vj = 0
					break
				}
				l = location{s.l.I + 1, s.l.J}
				if ('A' <= s.grid[l] && s.grid[l] <= 'Z') || s.grid[l] == '|' || s.grid[l] == '+' {
					// go down
					s.vi = 1
					s.vj = 0
					break
				}
				// we can't go anywhere -> panic
				panic(fmt.Sprintf("I can't move from location [%v]", s.l))
			}
		}
		if s.done {
			break
		}
	}
}

// Path http://adventofcode.com/2017/day/19
func Path(str string) string {
	s, err := newState(str)
	if err != nil {
		panic(err)
	}

	s.solve()

	return s.seen
}

// PathSteps http://adventofcode.com/2017/day/19
func PathSteps(str string) int {
	s, err := newState(str)
	if err != nil {
		panic(err)
	}

	s.solve()

	return s.steps - 1
}
