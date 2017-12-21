package advent2017

import (
	"fmt"
	"strconv"
	"strings"
)

type instruction struct {
	i, x, y string
}

func xAsInt(i instruction) (int, error) {
	x, err := strconv.Atoi(i.x)
	if err != nil {
		return 0, err
	}
	return x, nil
}

func yAsInt(i instruction) (int, error) {
	y, err := strconv.Atoi(i.y)
	if err != nil {
		return 0, err
	}
	return y, nil
}

type program2 struct {
	id           int
	instructions []instruction
	registers    map[string]int
	pos          int
	sends        int
}

func newProgram(id int, s string) *program2 {
	// parse instructions
	i := []instruction{}
	for _, l := range strings.Split(s, "\n") {
		t := strings.Split(l, " ")
		if len(t) == 2 {
			i = append(i, instruction{t[0], t[1], ""})
		} else {
			i = append(i, instruction{t[0], t[1], t[2]})
		}
	}

	r := make(map[string]int)
	r["p"] = id

	return &program2{id, i, r, 0, 0}
}

func peak(p *program2) string {
	return p.instructions[p.pos].i
}

func (p *program2) step(to, from chan int) {
	i := p.instructions[p.pos]
	switch i.i {
	case "snd":
		if x, err := xAsInt(i); err != nil {
			to <- p.registers[i.x]
		} else {
			to <- x
		}
		p.sends++
	case "set":
		if y, err := yAsInt(i); err != nil {
			p.registers[i.x] = p.registers[i.y]
		} else {
			p.registers[i.x] = y
		}
	case "add":
		if y, err := yAsInt(i); err != nil {
			p.registers[i.x] += p.registers[i.y]
		} else {
			p.registers[i.x] += y
		}
	case "mul":
		if y, err := yAsInt(i); err != nil {
			p.registers[i.x] *= p.registers[i.y]
		} else {
			p.registers[i.x] *= y
		}
	case "mod":
		if y, err := yAsInt(i); err != nil {
			p.registers[i.x] %= p.registers[i.y]
		} else {
			p.registers[i.x] %= y
		}
	case "rcv":
		p.registers[i.x] = <-from
	case "jgz":
		use := 0
		x, err := xAsInt(i)
		if err != nil {
			use = p.registers[i.x]
		} else {
			use = x
		}
		if use > 0 {
			if y, err := yAsInt(i); err != nil {
				p.pos += p.registers[i.y] - 1
			} else {
				p.pos += y - 1
			}
		}
	}

	p.pos++
}

// ParallelPrograms http://adventofcode.com/2017/day/18
func ParallelPrograms(s string) int {
	p0 := newProgram(0, s)
	p1 := newProgram(1, s)

	to0 := make(chan int, 1000)
	to1 := make(chan int, 1000)

	for count := 0; count < 100000; count++ {
		p0Blocked := peak(p0) == "rcv" && len(to0) == 0
		p1Blocked := peak(p1) == "rcv" && len(to1) == 0

		if p0Blocked && p1Blocked {
			return p1.sends
		}

		if !p0Blocked {
			p0.step(to1, to0)
		}

		if !p1Blocked {
			p1.step(to0, to1)
		}

		fmt.Println(p1.sends)
	}

	return p1.sends
}

// FirstRecover http://adventofcode.com/2017/day/18
// func FirstRecover(s string) int {
// 	instructions := []*instruction{}
// 	for _, l := range strings.Split(s, "\n") {
// 		t := strings.Split(l, " ")
// 		if len(t) == 2 {
// 			instructions = append(instructions, &instruction{t[0], t[1], ""})
// 		} else {
// 			instructions = append(instructions, &instruction{t[0], t[1], t[2]})
// 		}
// 	}

// 	registers := make(map[string]int)
// 	snd := 0
// 	i := 0
// 	done := false
// 	for {
// 		inst := instructions[i]
// 		switch inst.inst {
// 		case "snd":
// 			if x, err := inst.xAsInt(); err != nil {
// 				snd = registers[inst.x]
// 			} else {
// 				snd = x
// 			}
// 		case "set":
// 			if y, err := inst.yAsInt(); err != nil {
// 				registers[inst.x] = registers[inst.y]
// 			} else {
// 				registers[inst.x] = y
// 			}
// 		case "add":
// 			if y, err := inst.yAsInt(); err != nil {
// 				registers[inst.x] += registers[inst.y]
// 			} else {
// 				registers[inst.x] += y
// 			}
// 		case "mul":
// 			if y, err := inst.yAsInt(); err != nil {
// 				registers[inst.x] *= registers[inst.y]
// 			} else {
// 				registers[inst.x] *= y
// 			}
// 		case "mod":
// 			if y, err := inst.yAsInt(); err != nil {
// 				registers[inst.x] %= registers[inst.y]
// 			} else {
// 				registers[inst.x] %= y
// 			}
// 		case "rcv":
// 			if registers[inst.x] != 0 {
// 				done = true
// 			}
// 		case "jgz":
// 			if registers[inst.x] > 0 {
// 				if y, err := inst.yAsInt(); err != nil {
// 					i += registers[inst.y] - 1

// 				} else {
// 					i += y - 1
// 				}
// 			}
// 		}

// 		i++

// 		if done {
// 			break
// 		}
// 	}

// 	return snd
// }
