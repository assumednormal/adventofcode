package advent2017

import (
	"testing"
)

func TestDay01_Captcha_Part01(t *testing.T) {
	if v1 := SumSkipDigits("1122", 1); v1 != 3 {
		t.Error("Input 1122 got ", v1)
	}

	if v2 := SumSkipDigits("1111", 1); v2 != 4 {
		t.Error("Input 1111 got ", v2)
	}

	if v3 := SumSkipDigits("1234", 1); v3 != 0 {
		t.Error("Input 1234 got ", v3)
	}

	if v4 := SumSkipDigits("91212129", 1); v4 != 9 {
		t.Error("Input 91212129 got ", v4)
	}
}

func TestDay01_Captcha_Part02(t *testing.T) {
	if v1 := SumSkipDigits("1212", 2); v1 != 6 {
		t.Error("Input 1212 got ", v1)
	}

	if v2 := SumSkipDigits("1221", 2); v2 != 0 {
		t.Error("Input 1221 got ", v2)
	}

	if v3 := SumSkipDigits("123425", 3); v3 != 4 {
		t.Error("Input 123425 got ", v3)
	}

	if v4 := SumSkipDigits("123123", 3); v4 != 12 {
		t.Error("Input 123123 got ", v4)
	}

	if v5 := SumSkipDigits("12131415", 4); v5 != 4 {
		t.Error("Input 12131415 got ", v5)
	}
}

func TestDay02_Checksum_Part01(t *testing.T) {
	s := `5 1 9 5
7 5 3
2 4 6 8`
	if v := RowMaxDiffChecksum(s); v != 18 {
		t.Error("Input got ", v)
	}
}

func TestDay02_Checksum_Part02(t *testing.T) {
	s := `5 9 2 8
9 4 7 3
3 8 6 5`
	if v := RowDivChecksum(s); v != 9 {
		t.Error("Input got ", v)
	}
}

func TestDay03_Carry_Part01(t *testing.T) {
	if v1 := MovesToCenter(1); v1 != 0 {
		t.Error("Input 1 got ", v1)
	}

	if v2 := MovesToCenter(12); v2 != 3 {
		t.Error("Input 12 got ", v2)
	}

	if v3 := MovesToCenter(23); v3 != 2 {
		t.Error("Input 23 got ", v3)
	}

	if v4 := MovesToCenter(1024); v4 != 31 {
		t.Error("Input 1024 got ", v4)
	}
}

func TestDay03_Carry_Part02(t *testing.T) {
	if v1 := LeastCumulativeSpiral(1); v1 != 2 {
		t.Error("Input 1 got ", v1)
	}

	if v2 := LeastCumulativeSpiral(2); v2 != 4 {
		t.Error("Input 2 got ", v2)
	}

	if v3 := LeastCumulativeSpiral(3); v3 != 4 {
		t.Error("Input 3 got ", v3)
	}

	if v4 := LeastCumulativeSpiral(4); v4 != 5 {
		t.Error("Input 4 got ", v4)
	}

	if v5 := LeastCumulativeSpiral(5); v5 != 10 {
		t.Error("Input 5 got ", v5)
	}

	if v6 := LeastCumulativeSpiral(145); v6 != 147 {
		t.Error("Input 145 got ", v6)
	}
}

func TestDay04_Passphrases_Part01(t *testing.T) {
	if v1 := ValidPassphrases("aa bb cc dd ee"); v1 != 1 {
		t.Error("Input 1 got ", v1)
	}

	if v2 := ValidPassphrases("aa bb cc dd aa"); v2 != 0 {
		t.Error("Input 2 got ", v2)
	}

	if v3 := ValidPassphrases("aa bb cc dd aaa"); v3 != 1 {
		t.Error("Input 3 got ", v3)
	}
}

func TestDay04_Passphrases_Part02(t *testing.T) {
	if v1 := ValidAnagramPassphrases("abcde fghij"); v1 != 1 {
		t.Error("Input 1 got ", v1)
	}

	if v2 := ValidAnagramPassphrases("abcde xyz ecdab"); v2 != 0 {
		t.Error("Input 2 got ", v2)
	}

	if v3 := ValidAnagramPassphrases("a ab abc abd abf abj"); v3 != 1 {
		t.Error("Input 3 got ", v3)
	}

	if v4 := ValidAnagramPassphrases("iiii oiii ooii oooi oooo"); v4 != 1 {
		t.Error("Input 4 got ", v4)
	}

	if v5 := ValidAnagramPassphrases("oiii ioii iioi iiio"); v5 != 0 {
		t.Error("Input 5 got ", v5)
	}
}

func TestDay05_Steps_Part01(t *testing.T) {
	i := `0
3
0
1
-3`

	if v1 := StepsToExit(i); v1 != 5 {
		t.Error("Input 1 got ", v1)
	}
}

func TestDay05_Steps_Part02(t *testing.T) {
	i := `0
3
0
1
-3`

	if v1 := StepsToExit2(i); v1 != 10 {
		t.Error("Input 1 got ", v1)
	}
}

func TestDay06_Memory_Reallocation_Part01(t *testing.T) {
	i := "0 2 7 0"
	if v1 := StepsToPriorPattern(i); v1 != 5 {
		t.Error("Input 1 got ", v1)
	}
}

func TestDay06_Memory_Reallocation_Part02(t *testing.T) {
	i := "0 2 7 0"
	if v1 := StepsToPriorPattern2(i); v1 != 4 {
		t.Error("Input 1 got ", v1)
	}
}

func TestDay07_Recursive_Circus_Part01(t *testing.T) {
	i := `pbga (66)
xhth (57)
ebii (61)
havc (66)
ktlj (57)
fwft (72) -> ktlj, cntj, xhth
qoyq (66)
padx (45) -> pbga, havc, qoyq
tknk (41) -> ugml, padx, fwft
jptl (61)
ugml (68) -> gyxo, ebii, jptl
gyxo (61)
cntj (57)`

	if v1 := FindBottom(i); v1 != "tknk" {
		t.Error("Input 1 got ", v1)
	}
}

func TestDay07_Recursive_Circus_Part02(t *testing.T) {
	i := `pbga (66)
xhth (57)
ebii (61)
havc (66)
ktlj (57)
fwft (72) -> ktlj, cntj, xhth
qoyq (66)
padx (45) -> pbga, havc, qoyq
tknk (41) -> ugml, padx, fwft
jptl (61)
ugml (68) -> gyxo, ebii, jptl
gyxo (61)
cntj (57)`

	if v1 := BalanceTower(i); v1 != 60 {
		t.Error("Input 1 got ", v1)
	}
}

func TestDay08_Registers_Part01(t *testing.T) {
	i := `b inc 5 if a > 1
a inc 1 if b < 5
c dec -10 if a >= 1
c inc -20 if c == 10`

	if v1 := LargestRegister(i); v1 != 1 {
		t.Error("Input 1 got ", v1)
	}
}

func TestDay08_Registers_Part02(t *testing.T) {
	i := `b inc 5 if a > 1
a inc 1 if b < 5
c dec -10 if a >= 1
c inc -20 if c == 10`

	if v1 := LargestRegisterEver(i); v1 != 10 {
		t.Error("Input 1 got ", v1)
	}
}

func TestDay09_Streams_Part01(t *testing.T) {
	if v1 := StreamScore("{}"); v1 != 1 {
		t.Error("Input 1 got ", v1)
	}

	if v2 := StreamScore("{{{}}}"); v2 != 6 {
		t.Error("Input 2 got ", v2)
	}

	if v3 := StreamScore("{{},{}}"); v3 != 5 {
		t.Error("Input 3 got ", v3)
	}

	if v4 := StreamScore("{{{},{},{{}}}}"); v4 != 16 {
		t.Error("Input 4 got ", v4)
	}

	if v5 := StreamScore("{<a>,<a>,<a>,<a>}"); v5 != 1 {
		t.Error("Input 5 got ", v5)
	}

	if v6 := StreamScore("{{<ab>},{<ab>},{<ab>},{<ab>}}"); v6 != 9 {
		t.Error("Input 6 got ", v6)
	}

	if v7 := StreamScore("{{<!!>},{<!!>},{<!!>},{<!!>}}"); v7 != 9 {
		t.Error("Input 7 got ", v7)
	}

	if v8 := StreamScore("{{<a!>},{<a!>},{<a!>},{<ab>}}"); v8 != 3 {
		t.Error("Input 8 got ", v8)
	}
}

func TestDay09_Streams_Part02(t *testing.T) {
	if v1 := GarbageCounter("<>"); v1 != 0 {
		t.Error("Input 1 got ", v1)
	}

	if v2 := GarbageCounter("<random characters>"); v2 != 17 {
		t.Error("Input 2 got ", v2)
	}

	if v3 := GarbageCounter("<<<<>"); v3 != 3 {
		t.Error("Input 3 got ", v3)
	}

	if v4 := GarbageCounter("<{!>}>"); v4 != 2 {
		t.Error("Input 4 got ", v4)
	}

	if v5 := GarbageCounter("<!!>"); v5 != 0 {
		t.Error("Input 5 got ", v5)
	}

	if v6 := GarbageCounter("<!!!>>"); v6 != 0 {
		t.Error("Input 6 got ", v6)
	}

	if v7 := GarbageCounter(`<{o"i!a,<{i<a`); v7 != 10 {
		t.Error("Input 7 got ", v7)
	}
}

func TestDay10_Knot_Hash_Part01(t *testing.T) {
	if v1 := KnotHash(5, []int{3, 4, 1, 5}); v1 != 12 {
		t.Error("Input 1 got ", v1)
	}
}

func TestDay10_Knot_Hash_Part02(t *testing.T) {
	if v1 := KnotHash2([]byte("")); v1 != "a2582a3a0e66e6e86e3812dcb672a272" {
		t.Error("Input 1 got ", v1)
	}

	if v2 := KnotHash2([]byte("AoC 2017")); v2 != "33efeb34ea91902bb2f59c9920caa6cd" {
		t.Error("Input 2 got ", v2)
	}

	if v3 := KnotHash2([]byte("1,2,3")); v3 != "3efbe78a8d82f29979031a4aa0b16a9d" {
		t.Error("Input 3 got ", v3)
	}

	if v4 := KnotHash2([]byte("1,2,4")); v4 != "63960835bcdc130f0b66d7ff4f6a5a8e" {
		t.Error("Input 4 got ", v4)
	}
}

func TestDay11_Hex_Ed_Part01(t *testing.T) {
	if v1 := HexDistance("ne,ne,ne"); v1 != 3 {
		t.Error("Input 1 got ", v1)
	}

	if v2 := HexDistance("ne,ne,sw,sw"); v2 != 0 {
		t.Error("Input 2 got ", v2)
	}

	if v3 := HexDistance("ne,ne,s,s"); v3 != 2 {
		t.Error("Input 3 got ", v3)
	}

	if v4 := HexDistance("se,sw,se,sw,sw"); v4 != 3 {
		t.Error("Input 4 got ", v4)
	}
}

func TestDay11_Hex_Ed_Part02(t *testing.T) {
	if v1 := HexDistance2("ne,ne,ne"); v1 != 3 {
		t.Error("Input 1 got ", v1)
	}

	if v2 := HexDistance2("ne,ne,sw,sw"); v2 != 2 {
		t.Error("Input 2 got ", v2)
	}

	if v3 := HexDistance2("ne,ne,s,s"); v3 != 2 {
		t.Error("Input 3 got ", v3)
	}

	if v4 := HexDistance2("se,sw,se,sw,sw"); v4 != 3 {
		t.Error("Input 4 got ", v4)
	}
}

func TestDay12_Digital_Plumber_Part01(t *testing.T) {
	i := `0 <-> 2
1 <-> 1
2 <-> 0, 3, 4
3 <-> 2, 4
4 <-> 2, 3, 6
5 <-> 6
6 <-> 4, 5`

	if v1 := GroupWithZero(i); v1 != 6 {
		t.Error("Input 1 got ", v1)
	}
}

func TestDay12_Digital_Plumber_Part02(t *testing.T) {
	i := `0 <-> 2
1 <-> 1
2 <-> 0, 3, 4
3 <-> 2, 4
4 <-> 2, 3, 6
5 <-> 6
6 <-> 4, 5`

	if v1 := CountGroups(i); v1 != 2 {
		t.Error("Input 1 got ", v1)
	}
}

func TestDay13_Packet_Scanner_Part01(t *testing.T) {
	i := `0: 3
1: 2
4: 4
6: 4`

	if v1 := TripSeverity(i); v1 != 24 {
		t.Error("Input 1 got ", v1)
	}
}

func TestDay13_Packet_Scanner_Part02(t *testing.T) {
	i := `0: 3
1: 2
4: 4
6: 4`

	if v1 := MinimumDelay(i); v1 != 10 {
		t.Error("Input 1 got ", v1)
	}
}

func TestDay14_Disk_Defragmentation_Part01(t *testing.T) {
	if v1 := UsedSquares("flqrgnkx"); v1 != 8108 {
		t.Error("Input 1 got ", v1)
	}
}

func TestDay14_Disk_Defragmentation_Part02(t *testing.T) {
	if v1 := CountRegions("flqrgnkx"); v1 != 1242 {
		t.Error("Input 1 got ", v1)
	}
}

func TestDay15_Dueling_Generators_Part01(t *testing.T) {
	if v1 := MatchingBits(65, 16807, 8921, 48271, 5); v1 != 1 {
		t.Error("Input 1 got ", v1)
	}

	if v2 := MatchingBits(65, 16807, 8921, 48271, 40000000); v2 != 588 {
		t.Error("Input 2 got ", v2)
	}
}

func TestDay15_Dueling_Generators_Part02(t *testing.T) {
	if v1 := MatchingBits2(65, 16807, 4, 8921, 48271, 8, 5000000); v1 != 309 {
		t.Error("Input 1 got ", v1)
	}
}

func TestDay16_Permutation_Promenade_Part01(t *testing.T) {
	if v1 := Dance("s1,x3/4,pe/b", []rune("abcde")); v1 != "baedc" {
		t.Error("Input 1 got ", v1)
	}
}

func TestDay17_Spinlock_Part01(t *testing.T) {
	if v1 := Spinlock(9, 3); v1 != 5 {
		t.Error("Input 1 got ", v1)
	}

	if v2 := Spinlock(2017, 3); v2 != 638 {
		t.Error("Input 2 got ", v2)
	}
}

func TestDay17_Spinlock_Part02(t *testing.T) {
	if v1 := Spinlock2(9, 3); v1 != 9 {
		t.Error("Input 1 got ", v1)
	}

	if v2 := Spinlock2(2017, 3); v2 != 1226 {
		t.Error("Input 2 got ", v2)
	}
}

// func TestDay18_Duet_Part01(t *testing.T) {
// 	i := `set a 1
// add a 2
// mul a a
// mod a 5
// snd a
// set a 0
// rcv a
// jgz a -1
// set a 1
// jgz a -2`

// 	if v1 := FirstRecover(i); v1 != 4 {
// 		t.Error("Input 1 got ", v1)
// 	}
// }

func TestDay18_Duet_Part02(t *testing.T) {
	i := `snd 1
snd 2
snd p
rcv a
rcv b
rcv c
rcv d`

	if v1 := ParallelPrograms(i); v1 != 3 {
		t.Error("Input 1 got ", v1)
	}
}

func TestDay18_Duet_Part03(t *testing.T) {
	i := `rcv a`

	if v1 := ParallelPrograms(i); v1 != 0 {
		t.Error("Input 1 got ", v1)
	}
}
