package advent2017

import "testing"

func TestDay01_Captcha(t *testing.T) {
	if v1 := SumConsecutiveDigits("1122"); v1 != 3 {
		t.Error("Input 1122 got ", v1)
	}

	if v2 := SumConsecutiveDigits("1111"); v2 != 4 {
		t.Error("Input 1111 got ", v2)
	}

	if v3 := SumConsecutiveDigits("1234"); v3 != 0 {
		t.Error("Input 1234 got ", v3)
	}

	if v4 := SumConsecutiveDigits("91212129"); v4 != 9 {
		t.Error("Input 91212129 got ", v4)
	}
}
