package advent2017

import (
	"fmt"
	"strings"
)

type program3 struct {
	instructions []instruction
	registers    map[string]int
	pos          int
	muls         int
	done         bool
}

func newProgram3(s string, a int) *program3 {
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
	r["a"] = a

	return &program3{i, r, 0, 0, false}
}

func (p *program3) String() string {
	return fmt.Sprintf("a=%d,b=%d,c=%d,d=%d,e=%d,f=%d,g=%d,h=%d",
		p.registers["a"], p.registers["b"], p.registers["c"], p.registers["d"],
		p.registers["e"], p.registers["f"], p.registers["g"], p.registers["h"])
}

func (p *program3) step() {
	if p.done {
		return
	}

	if p.pos < 0 || p.pos > len(p.instructions)-1 {
		p.done = true
		return
	}

	i := p.instructions[p.pos]
	switch i.i {
	case "set":
		if y, err := yAsInt(i); err != nil {
			p.registers[i.x] = p.registers[i.y]
		} else {
			p.registers[i.x] = y
		}
	case "sub":
		if y, err := yAsInt(i); err != nil {
			p.registers[i.x] -= p.registers[i.y]
		} else {
			p.registers[i.x] -= y
		}
	case "mul":
		if y, err := yAsInt(i); err != nil {
			p.registers[i.x] *= p.registers[i.y]
		} else {
			p.registers[i.x] *= y
		}
		p.muls++
	case "jnz":
		use := 0
		x, err := xAsInt(i)
		if err != nil {
			use = p.registers[i.x]
		} else {
			use = x
		}
		if use != 0 {
			if y, err := yAsInt(i); err != nil {
				p.pos += p.registers[i.y] - 1
			} else {
				p.pos += y - 1
			}
		}
	}

	p.pos++
}

// CountMultiplies http://adventofcode.com/2017/day/23
func CountMultiplies(s string) int {
	p := newProgram3(s, 0)

	for !p.done {
		p.step()
	}

	return p.muls
}

// RegisterH http://adventofcode.com/2017/day/23
func RegisterH(s string) int {
	p := newProgram3(s, 1)

	for !p.done {
		p.step()
	}

	return p.registers["h"]
}

func run() int {
	var b, c, d, f, g, h int

	b = 106700
	c = 123700

	for {
		f = 1
		d = 2
		for {
			if b%d == 0 {
				f = 0
			}
			// e = 2
			// for e = 2; e <= b; e++ {
			// 	// determine if d is a divisor of b?
			// 	// if d*e == b, set f to 0 (indicating b is not prime?)
			// 	// iterate until e == b (this means d*e overshoots b massively)
			// 	g = d*e - b
			// 	if g == 0 {
			// 		f = 0
			// 	}
			// 	e++
			// 	g = e - b
			// 	if g == 0 {
			// 		break
			// 	}
			// }
			d++
			g = d - b
			if g == 0 {
				break
			}
		}
		if g != 0 {
			continue
		}
		if f == 0 {
			h++
		}
		g = b - c
		if g == 0 {
			break
		}
		b += 17
	}

	return h
}
