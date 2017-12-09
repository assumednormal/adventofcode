package main

import (
	"fmt"
	"io/ioutil"

	"adventofcode2017/advent2017"
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
	day01Input := readFile("day01_input.txt")
	fmt.Println(advent2017.SumSkipDigits(day01Input, 1))
	fmt.Println(advent2017.SumSkipDigits(day01Input, len(day01Input)/2))

	// day 2
	day02Input := readFile("day02_input.txt")
	fmt.Println(advent2017.RowMaxDiffChecksum(day02Input))
	fmt.Println(advent2017.RowDivChecksum(day02Input))

	// day 3
	fmt.Println(advent2017.MovesToCenter(325489))
	fmt.Println(advent2017.LeastCumulativeSpiral(325489))

	// day 4
	day04Input := readFile("day04_input.txt")
	fmt.Println(advent2017.ValidPassphrases(day04Input))
	fmt.Println(advent2017.ValidAnagramPassphrases(day04Input))

	// day 5
	day05Input := readFile("day05_input.txt")
	fmt.Println(advent2017.StepsToExit(day05Input))
	fmt.Println(advent2017.StepsToExit2(day05Input))
}
