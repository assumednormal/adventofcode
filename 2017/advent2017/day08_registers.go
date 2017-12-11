package advent2017

import (
	"strconv"
	"strings"
)

// LargestRegister http://adventofcode.com/2017/day/8
func LargestRegister(s string) int {
	registers := make(map[string]int)
	for _, l := range strings.Split(s, "\n") {
		f := strings.Fields(l)

		register := f[0]
		operation := f[1]
		delta, err := strconv.Atoi(f[2])
		if err != nil {
			panic(err)
		}
		condRegister := f[4]
		condOperation := f[5]
		condValue, err := strconv.Atoi(f[6])
		if err != nil {
			panic(err)
		}

		// always increment
		if operation == "dec" {
			delta = -delta
		}

		switch condOperation {
		case "!=":
			if registers[condRegister] != condValue {
				registers[register] += delta
			}
		case "<":
			if registers[condRegister] < condValue {
				registers[register] += delta
			}
		case "<=":
			if registers[condRegister] <= condValue {
				registers[register] += delta
			}
		case "==":
			if registers[condRegister] == condValue {
				registers[register] += delta
			}
		case ">=":
			if registers[condRegister] >= condValue {
				registers[register] += delta
			}
		case ">":
			if registers[condRegister] > condValue {
				registers[register] += delta
			}
		}
	}

	max := 0
	for _, v := range registers {
		if max < v {
			max = v
		}
	}

	return max
}

// LargestRegisterEver http://adventofcode.com/2017/day/8
func LargestRegisterEver(s string) int {
	registers := make(map[string]int)
	max := 0
	for _, l := range strings.Split(s, "\n") {
		f := strings.Fields(l)

		register := f[0]
		operation := f[1]
		delta, err := strconv.Atoi(f[2])
		if err != nil {
			panic(err)
		}
		condRegister := f[4]
		condOperation := f[5]
		condValue, err := strconv.Atoi(f[6])
		if err != nil {
			panic(err)
		}

		// always increment
		if operation == "dec" {
			delta = -delta
		}

		switch condOperation {
		case "!=":
			if registers[condRegister] != condValue {
				registers[register] += delta
			}
		case "<":
			if registers[condRegister] < condValue {
				registers[register] += delta
			}
		case "<=":
			if registers[condRegister] <= condValue {
				registers[register] += delta
			}
		case "==":
			if registers[condRegister] == condValue {
				registers[register] += delta
			}
		case ">=":
			if registers[condRegister] >= condValue {
				registers[register] += delta
			}
		case ">":
			if registers[condRegister] > condValue {
				registers[register] += delta
			}
		}

		if max < registers[register] {
			max = registers[register]
		}
	}

	return max
}
