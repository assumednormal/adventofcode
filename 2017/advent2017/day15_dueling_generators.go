package advent2017

import "math"

type generator struct {
	seed, factor, value, mod uint
}

func newGenerator(seed, factor, mod uint) *generator {
	return &generator{seed, factor, (seed * factor) % uint(math.MaxInt32), mod}
}

func (g *generator) step() uint {
	for {
		g.value = (g.value * g.factor) % uint(math.MaxInt32)
		if g.value%g.mod == 0 {
			break
		}
	}
	return g.value
}

// MatchingBits http://adventofcode.com/2017/day/15
func MatchingBits(aSeed, aFactor, bSeed, bFactor, rounds uint) uint {
	a := newGenerator(aSeed, aFactor, 1)
	b := newGenerator(bSeed, bFactor, 1)

	matching := uint(0)
	for i := uint(0); i < rounds; i++ {
		aVal := a.step()
		bVal := b.step()

		if aVal&0xFFFF == bVal&0xFFFF {
			matching++
		}
	}

	return matching
}

// MatchingBits2 http://adventofcode.com/2017/day/15
func MatchingBits2(aSeed, aFactor, aMod, bSeed, bFactor, bMod, rounds uint) uint {
	a := newGenerator(aSeed, aFactor, aMod)
	b := newGenerator(bSeed, bFactor, bMod)

	matching := uint(0)
	for i := uint(0); i < rounds; i++ {
		aVal := a.step()
		bVal := b.step()

		if aVal&0xFFFF == bVal&0xFFFF {
			matching++
		}
	}

	return matching
}
