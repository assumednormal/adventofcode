package advent2017

func absDiff(a, b int) int {
	if a < b {
		return b - a
	}
	return a - b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// MovesToCenter http://adventofcode.com/2017/day/3
func MovesToCenter(v int) int {
	ring := 0
	for i := 0; (2*i+1)*(2*i+1) < v; i++ {
		ring++
	}
	sideLength := 2*ring + 1
	minDist := sideLength - 1
	for i := (sideLength-2)*(sideLength-2) + (sideLength-1)/2; i <= (sideLength)*(sideLength) && sideLength > 1; i += sideLength - 1 {
		minDist = min(minDist, ring+absDiff(i, v))
	}
	return minDist
}

func calculateValue(x, y int, m map[int]map[int]int) int {
	v := 0
	if _, ok := m[x-1]; ok {
		v += m[x-1][y-1] + m[x-1][y] + m[x-1][y+1]
	}
	if _, ok := m[x]; ok {
		v += m[x][y-1] + m[x][y+1]
	}
	if _, ok := m[x+1]; ok {
		v += m[x+1][y-1] + m[x+1][y] + m[x+1][y+1]
	}
	return v
}

func checkValue(p, q int) bool {
	if p < q {
		return true
	}
	return false
}

// LeastCumulativeSpiral http://adventofcode.com/2017/day/3
func LeastCumulativeSpiral(p int) int {
	// keep track of existing values in a map
	vals := make(map[int]map[int]int)
	vals[0] = make(map[int]int)
	vals[0][0] = 1

	x := 0
	y := 0
	for sideLength := 3; true; sideLength += 2 {
		// move right once and down once
		x++
		y++

		// calculate values going up right side
		if _, ok := vals[x]; !ok {
			vals[x] = make(map[int]int)
		}
		for i := 0; i < sideLength-1; i++ {
			y--
			vals[x][y] = calculateValue(x, y, vals)
			if checkValue(p, vals[x][y]) {
				return vals[x][y]
			}
		}

		// calculate values going across top side
		for i := 0; i < sideLength-1; i++ {
			x--
			if _, ok := vals[x]; !ok {
				vals[x] = make(map[int]int)
			}
			vals[x][y] = calculateValue(x, y, vals)
			if checkValue(p, vals[x][y]) {
				return vals[x][y]
			}
		}

		// calculate values going down left side
		if _, ok := vals[x]; !ok {
			vals[x] = make(map[int]int)
		}
		for i := 0; i < sideLength-1; i++ {
			y++
			vals[x][y] = calculateValue(x, y, vals)
			if checkValue(p, vals[x][y]) {
				return vals[x][y]
			}
		}

		// calculate values going across bottom side
		for i := 0; i < sideLength-1; i++ {
			x++
			if _, ok := vals[x]; !ok {
				vals[x] = make(map[int]int)
			}
			vals[x][y] = calculateValue(x, y, vals)
			if checkValue(p, vals[x][y]) {
				return vals[x][y]
			}
		}
	}

	return -1
}
