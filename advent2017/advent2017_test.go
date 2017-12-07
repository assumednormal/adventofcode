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
