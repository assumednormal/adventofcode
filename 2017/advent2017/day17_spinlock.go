package advent2017

import "fmt"

type circularBuffer struct {
	buf  []int
	step int
	cur  int
}

func newCircularBuffer(step int) *circularBuffer {
	return &circularBuffer{[]int{0}, step, 0}
}

func (b *circularBuffer) insert(v int) int {
	if len(b.buf) == 1 {
		b.cur = 1
		b.buf = append(b.buf, v)
		return b.buf[(b.cur+1)%len(b.buf)]
	}
	b.cur = (b.cur + b.step + 1) % len(b.buf)
	tmp := append([]int{v}, b.buf[b.cur:]...)
	b.buf = append(b.buf[:b.cur], tmp...)
	return b.buf[(b.cur+1)%len(b.buf)]
}

func (b *circularBuffer) valAfter(v int) int {
	var i int
	for i = 0; i < len(b.buf); i++ {
		if b.buf[i] == v {
			break
		}
	}
	return b.buf[(i+1)%len(b.buf)]
}

func (b *circularBuffer) String() string {
	return fmt.Sprintf("Current Position: %d\nStep Size: %d\nBuffer: %v\n", b.cur, b.step, b.buf)
}

// Spinlock http://adventofcode.com/2017/day/17
func Spinlock(maxVal, step int) int {
	b := newCircularBuffer(step)
	var v int
	for i := 1; i <= maxVal; i++ {
		v = b.insert(i)
	}
	return v
}

// Spinlock2 http://adventofcode.com/2017/day/17
func Spinlock2(maxVal, step int) int {
	length := 1
	pos := 0
	second := 0

	for i := 1; i <= maxVal; i++ {
		pos = (pos + step) % length
		if pos == 0 {
			second = i
		}
		length++
		pos++
	}

	return second
}
