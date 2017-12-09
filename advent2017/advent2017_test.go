package advent2017

import "testing"

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
