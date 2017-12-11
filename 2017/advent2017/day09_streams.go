package advent2017

const (
	leftCurlyBracket  = '{' // beginning of group
	rightCurlyBracket = '}' // end of group
	leftAngleBracket  = '<' // beginning of garbage
	rightAngleBracket = '>' // end of garbage
	bang              = '!' // ignore next character
)

// StreamScore http://adventofcode.com/2017/day/9
func StreamScore(s string) int {
	totalScore := 0

	depth := 0
	garbage := false
	ignore := false

	for _, c := range s {
		switch {
		case ignore:
			ignore = false
		case c == bang:
			ignore = true
		case c == rightAngleBracket:
			garbage = false
		case c == leftAngleBracket:
			garbage = true
		case garbage:
		case c == rightCurlyBracket:
			totalScore += depth
			depth--
		case c == leftCurlyBracket:
			depth++
		}
	}

	return totalScore
}

// GarbageCounter http://adventofcode.com/2017/day/9
func GarbageCounter(s string) int {
	totalGarbage := 0

	garbage := false
	ignore := false

	for _, c := range s {
		switch {
		case ignore:
			ignore = false
		case c == bang:
			ignore = true
		case c == rightAngleBracket:
			garbage = false
		case garbage:
			totalGarbage++
		case c == leftAngleBracket:
			garbage = true
		}
	}

	return totalGarbage
}
