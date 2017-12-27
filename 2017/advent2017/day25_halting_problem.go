package advent2017

type haltingProblem struct {
	tape     map[int]bool
	position int
	state    rune
}

func newHaltingProblem(state rune) *haltingProblem {
	return &haltingProblem{make(map[int]bool), 0, state}
}

func (hp *haltingProblem) step() {
	switch hp.state {
	case 'A':
		if !hp.tape[hp.position] {
			hp.tape[hp.position] = true
			hp.position++
			hp.state = 'B'
		} else {
			delete(hp.tape, hp.position)
			hp.position++
			hp.state = 'C'
		}
	case 'B':
		if !hp.tape[hp.position] {
			delete(hp.tape, hp.position)
			hp.position--
			hp.state = 'A'
		} else {
			delete(hp.tape, hp.position)
			hp.position++
			hp.state = 'D'
		}
	case 'C':
		if !hp.tape[hp.position] {
			hp.tape[hp.position] = true
			hp.position++
			hp.state = 'D'
		} else {
			hp.tape[hp.position] = true
			hp.position++
			hp.state = 'A'
		}
	case 'D':
		if !hp.tape[hp.position] {
			hp.tape[hp.position] = true
			hp.position--
			hp.state = 'E'
		} else {
			delete(hp.tape, hp.position)
			hp.position--
			hp.state = 'D'
		}
	case 'E':
		if !hp.tape[hp.position] {
			hp.tape[hp.position] = true
			hp.position++
			hp.state = 'F'
		} else {
			hp.tape[hp.position] = true
			hp.position--
			hp.state = 'B'
		}
	case 'F':
		if !hp.tape[hp.position] {
			hp.tape[hp.position] = true
			hp.position++
			hp.state = 'A'
		} else {
			hp.tape[hp.position] = true
			hp.position++
			hp.state = 'E'
		}
	}
}

func (hp *haltingProblem) checksum() int {
	s := 0
	for _, v := range hp.tape {
		if v {
			s++
		}
	}
	return s
}

// HaltingProblem https://adventofcode.com/2017/day/25
func HaltingProblem(steps int, state rune) int {
	hp := newHaltingProblem(state)
	for i := 0; i < steps; i++ {
		hp.step()
	}
	return hp.checksum()
}
