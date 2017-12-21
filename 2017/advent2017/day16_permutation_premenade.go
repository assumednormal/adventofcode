package advent2017

import (
	"fmt"
	"strconv"
	"strings"
)

// Dance http://adventofcode.com/2017/day/16
func Dance(s string, programs []rune) string {
	for _, m := range strings.Split(s, ",") {
		switch m[0] {
		case 's':
			n, err := strconv.Atoi(m[1:])
			if err != nil {
				panic(err)
			}
			for i := 0; i < n; i++ {
				programs = append([]rune{programs[len(programs)-1]}, programs[0:len(programs)-1]...)
			}
		case 'x':
			q := strings.Split(m[1:], "/")
			a, err := strconv.Atoi(q[0])
			if err != nil {
				panic(err)
			}
			b, err := strconv.Atoi(q[1])
			if err != nil {
				panic(err)
			}
			programs[a], programs[b] = programs[b], programs[a]
		case 'p':
			q := strings.Split(m[1:], "/")
			var a, b int
			for i := 0; i < len(programs); i++ {
				if programs[i] == []rune(q[0])[0] {
					a = i
				} else if programs[i] == []rune(q[1])[0] {
					b = i
				}
			}
			programs[a], programs[b] = programs[b], programs[a]
		}
	}
	return string(programs)
}

// Dance2 http://adventofcode.com/2017/day/16
func Dance2(s string, programs []rune) string {
	programOrders := make(map[string]int)
	prevProgram := string(programs)
	programOrders[prevProgram] = 0
	var i int
	for i = 0; i < 1000000000; i++ {
		prevProgram = Dance(s, []rune(prevProgram))
		if _, ok := programOrders[prevProgram]; ok {
			break
		} else {
			programOrders[prevProgram] = i + 1
		}
	}
	cycle := i + 1
	remainingCycles := 1000000000 % cycle
	fmt.Println(cycle, remainingCycles)
	for k, v := range programOrders {
		if v == remainingCycles {
			return k
		}
	}
	return "not found"
}
