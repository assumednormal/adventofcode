package advent2017

import (
	"strconv"
	"strings"
)

type program struct {
	Name     string
	Weight   int
	Children []string
}

// FindBottom http://adventofcode.com/2017/day/7
func FindBottom(s string) string {
	// build list of nodes + children
	programs := make(map[string]program)
	for _, l := range strings.Split(s, "\n") {
		f := strings.Fields(l)
		name := f[0]
		weight, err := strconv.Atoi(f[1][1:(len(f[1]) - 1)])
		if err != nil {
			panic(err)
		}
		children := []string{}
		if len(f) > 2 {
			for _, child := range f[3:] {
				children = append(children, strings.Replace(child, ",", "", -1))
			}
		}
		programs[name] = program{name, weight, children}
	}

	// which nodes have children
	bottoms := make(map[string]bool)
	for _, p := range programs {
		if len(p.Children) > 0 {
			bottoms[p.Name] = true
		}
	}

	// remove nodes that are children
	for _, p := range programs {
		for _, c := range p.Children {
			delete(bottoms, c)
		}
	}

	for k := range bottoms {
		return k
	}

	return ""
}

func totalWeight(name string, programs map[string]program) int {
	w := programs[name].Weight
	for _, c := range programs[name].Children {
		w += totalWeight(c, programs)
	}
	return w
}

// BalanceTower http://adventofcode.com/2017/day/7
func BalanceTower(s string) int {
	// build list of nodes + children
	programs := make(map[string]program)
	for _, l := range strings.Split(s, "\n") {
		f := strings.Fields(l)
		name := f[0]
		weight, err := strconv.Atoi(f[1][1:(len(f[1]) - 1)])
		if err != nil {
			panic(err)
		}
		children := []string{}
		if len(f) > 2 {
			for _, child := range f[3:] {
				children = append(children, strings.Replace(child, ",", "", -1))
			}
		}
		programs[name] = program{name, weight, children}
	}

	root := FindBottom(s)
	d := 0
	for {
		weights := make(map[int][]string)
		for _, c := range programs[root].Children {
			tw := totalWeight(c, programs)
			if _, ok := weights[tw]; ok {
				weights[tw] = append(weights[tw], c)
			} else {
				weights[tw] = []string{c}
			}
		}
		if len(weights) == 1 {
			break
		}
		oddWeight := 0
		oddBase := 0
		oddProgram := ""
		commonWeight := 0
		for k, v := range weights {
			if len(v) == 1 {
				oddWeight = k
				oddBase = programs[v[0]].Weight
				oddProgram = v[0]
			} else {
				commonWeight = k
			}
		}
		d = commonWeight - (oddWeight - oddBase)
		root = oddProgram
	}

	return d
}
