package advent2017

import "strconv"

// KnotHash http://adventofcode.com/2017/day/10
func KnotHash(n int, lengths []int) int {
	list := make([]int, n)
	for i := 0; i < n; i++ {
		list[i] = i
	}

	p := 0
	for i, l := range lengths {
		for j := 0; j <= (l-1)/2; j++ {
			list[(p+j)%n], list[(p+l-1-j)%n] = list[(p+l-1-j)%n], list[(p+j)%n]
		}
		p = (p + l + i) % n
	}

	return list[0] * list[1]
}

// KnotHash2 http://adventofcode.com/2017/day/10
func KnotHash2(lengths []byte) string {
	lengths = append(lengths, []byte{17, 31, 73, 47, 23}...)

	list := make([]int, 256)
	for i := 0; i < 256; i++ {
		list[i] = i
	}

	p := 0
	skip := 0
	for a := 0; a < 64; a++ {
		for _, l := range lengths {
			for j := 0; j <= (int(l)-1)/2; j++ {
				list[(p+j)%256], list[(p+int(l)-1-j)%256] = list[(p+int(l)-1-j)%256], list[(p+j)%256]
			}
			p = (p + int(l) + skip) % 256
			skip++
		}
	}

	dense := make([]int, 16)
	for i := 0; i < 16; i++ {
		dense[i] = list[16*i]
		for j := 0; j < 15; j++ {
			dense[i] ^= list[16*i+j+1]
		}
	}

	hex := ""
	for i := 0; i < 16; i++ {
		add := strconv.FormatInt(int64(dense[i]), 16)
		if len(add) == 1 {
			hex += "0"
		}
		hex += add
	}

	return hex
}
