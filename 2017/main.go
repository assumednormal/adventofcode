package main

import (
	"io/ioutil"
)

func readFile(f string) string {
	b, err := ioutil.ReadFile(f)
	if err != nil {
		panic(err)
	}
	return string(b)
}

func main() {
	// day 1
	// day01Input := readFile("day01_input.txt")
	// fmt.Println(advent2017.SumSkipDigits(day01Input, 1))
	// fmt.Println(advent2017.SumSkipDigits(day01Input, len(day01Input)/2))

	// day 2
	// day02Input := readFile("day02_input.txt")
	// fmt.Println(advent2017.RowMaxDiffChecksum(day02Input))
	// fmt.Println(advent2017.RowDivChecksum(day02Input))

	// day 3
	// fmt.Println(advent2017.MovesToCenter(325489))
	// fmt.Println(advent2017.LeastCumulativeSpiral(325489))

	// day 4
	// day04Input := readFile("day04_input.txt")
	// fmt.Println(advent2017.ValidPassphrases(day04Input))
	// fmt.Println(advent2017.ValidAnagramPassphrases(day04Input))

	// day 5
	// day05Input := readFile("day05_input.txt")
	// fmt.Println(advent2017.StepsToExit(day05Input))
	// fmt.Println(advent2017.StepsToExit2(day05Input))

	// day 6
	// day06Input := readFile("day06_input.txt")
	// fmt.Println(advent2017.StepsToPriorPattern(day06Input))
	// fmt.Println(advent2017.StepsToPriorPattern2(day06Input))

	// day 7
	// day07Input := readFile("day07_input.txt")
	// fmt.Println(advent2017.FindBottom(day07Input))
	// fmt.Println(advent2017.BalanceTower(day07Input))

	// day 8
	// day08Input := readFile("day08_input.txt")
	// fmt.Println(advent2017.LargestRegister(day08Input))
	// fmt.Println(advent2017.LargestRegisterEver(day08Input))

	// day 9
	// day09Input := readFile("day09_input.txt")
	// fmt.Println(advent2017.StreamScore(day09Input))
	// fmt.Println(advent2017.GarbageCounter(day09Input))

	// day 10
	// fmt.Println(advent2017.KnotHash(256, []int{14, 58, 0, 116, 179, 16, 1, 104, 2, 254, 167, 86, 255, 55, 122, 244}))
	// fmt.Println(advent2017.KnotHash2([]byte("14,58,0,116,179,16,1,104,2,254,167,86,255,55,122,244")))

	// day 11
	// day11Input := readFile("day11_input.txt")
	// fmt.Println(advent2017.HexDistance(day11Input))
	// fmt.Println(advent2017.HexDistance2(day11Input))

	// day 12
	// day12Input := readFile("day12_input.txt")
	// fmt.Println(advent2017.GroupWithZero(day12Input))
	// fmt.Println(advent2017.CountGroups(day12Input))

	// day 13
	// day13Input := readFile("day13_input.txt")
	// fmt.Println(advent2017.TripSeverity(day13Input))
	// fmt.Println(advent2017.MinimumDelay(day13Input))

	// day 14
	// fmt.Println(advent2017.UsedSquares("oundnydw"))
	// fmt.Println(advent2017.CountRegions("oundnydw"))

	// day 15
	// fmt.Println(advent2017.MatchingBits(289, 16807, 629, 48271, 40000000))
	// fmt.Println(advent2017.MatchingBits2(289, 16807, 4, 629, 48271, 8, 5000000))

	// day 16
	// day16Input := readFile("day16_input.txt")
	// fmt.Println(advent2017.Dance(day16Input, []rune("abcdefghijklmnop")))
	// fmt.Println(advent2017.Dance2(day16Input, []rune("abcdefghijklmnop")))

	// day 17
	// fmt.Println(advent2017.Spinlock(2017, 316))
	// fmt.Println(advent2017.Spinlock2(50000000, 316))

	// day 18
	// day18Input := readFile("day18_input.txt")
	// fmt.Println(advent2017.FirstRecover(day18Input))
	// fmt.Println(advent2017.ParallelPrograms(day18Input))

	// day 19
	// day19Input := readFile("day19_input.txt")
	// fmt.Println(advent2017.Path(day19Input))
	// fmt.Println(advent2017.PathSteps(day19Input))

	// day 20
	// day20Input := readFile("day20_input.txt")
	// fmt.Println(advent2017.ClosestParticle(day20Input))
	// fmt.Println(advent2017.ParticlesAfterCollisions(day20Input))

	// day 22
	// day22Input := readFile("day22_input.txt")
	// fmt.Println(advent2017.Infections(day22Input, 10000))
	// fmt.Println(advent2017.Infections2(day22Input, 10000000))

	// day 23
	// day23Input := readFile("day23_input.txt")
	// fmt.Println(advent2017.CountMultiplies(day23Input))
	// fmt.Println(advent2017.RegisterH(day23Input))
	// day 23 is solved in a test...
}
