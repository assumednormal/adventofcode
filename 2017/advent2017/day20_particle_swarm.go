package advent2017

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

type particle struct {
	x, y, z, vx, vy, vz, ax, ay, az int
}

func newParticle(x, y, z, vx, vy, vz, ax, ay, az int) *particle {
	return &particle{x, y, z, vx, vy, vz, ax, ay, az}
}

func (p *particle) step() {
	// velocity
	p.vx += p.ax
	p.vy += p.ay
	p.vz += p.az

	// position
	p.x += p.vx
	p.y += p.vy
	p.z += p.vz
}

func abs(a int) int {
	if a > 0 {
		return a
	}
	return -a
}

func (p *particle) manhattanDist(x, y, z int) int {
	return abs(p.x-x) + abs(p.y-y) + abs(p.z-z)
}

func intToStr(s string) int {
	t, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return t
}

// ClosestParticle http://adventofcode.com/2017/day/20
func ClosestParticle(s string) int {
	particles := []*particle{}
	for _, l := range strings.Split(s, "\n") {
		// p=< 3,0,0>, v=< 2,0,0>, a=<-1,0,0>
		pva := strings.Split(strings.Replace(strings.Replace(strings.Replace(strings.Replace(strings.Replace(l, "p=<", "", 1), "v=<", "", 1), "a=<", "", 1), " ", "", -1), ">", "", -1), ",")
		x := intToStr(pva[0])
		y := intToStr(pva[1])
		z := intToStr(pva[2])
		vx := intToStr(pva[3])
		vy := intToStr(pva[4])
		vz := intToStr(pva[5])
		ax := intToStr(pva[6])
		ay := intToStr(pva[7])
		az := intToStr(pva[8])
		particles = append(particles, newParticle(x, y, z, vx, vy, vz, ax, ay, az))
	}

	closest := 0
	for j := 1; j < len(particles); j++ {
		if abs(particles[closest].ax)+abs(particles[closest].ay)+abs(particles[closest].az) >
			abs(particles[j].ax)+abs(particles[j].ay)+abs(particles[j].az) {
			closest = j
		}
	}

	return closest
}

// ParticlesAfterCollisions http://adventofcode.com/2017/day/20
func ParticlesAfterCollisions(s string) int {
	particles := []*particle{}
	for _, l := range strings.Split(s, "\n") {
		// p=< 3,0,0>, v=< 2,0,0>, a=<-1,0,0>
		pva := strings.Split(strings.Replace(strings.Replace(strings.Replace(strings.Replace(strings.Replace(l, "p=<", "", 1), "v=<", "", 1), "a=<", "", 1), " ", "", -1), ">", "", -1), ",")
		x := intToStr(pva[0])
		y := intToStr(pva[1])
		z := intToStr(pva[2])
		vx := intToStr(pva[3])
		vy := intToStr(pva[4])
		vz := intToStr(pva[5])
		ax := intToStr(pva[6])
		ay := intToStr(pva[7])
		az := intToStr(pva[8])
		particles = append(particles, newParticle(x, y, z, vx, vy, vz, ax, ay, az))
	}

	for k := 1; k < 100000; k++ {
		// remove collisions
		toRemove := make(map[int]bool)
		for i := len(particles) - 1; i >= 0; i-- {
			for j := i - 1; j >= 0; j-- {
				if particles[i].manhattanDist(particles[j].x, particles[j].y, particles[j].z) == 0 {
					toRemove[i] = true
					toRemove[j] = true
				}
			}
		}

		toRemoveS := []int{}
		for k := range toRemove {
			toRemoveS = append(toRemoveS, k)
		}

		sort.Sort(sort.Reverse(sort.IntSlice(toRemoveS)))

		for _, idx := range toRemoveS {
			particles = append(particles[:idx], particles[idx+1:]...)
		}

		// step particles
		for _, p := range particles {
			p.step()
		}

		if k%1000 == 0 {
			fmt.Println("--- Iteration", k, "has", len(particles), "particles")
		}
	}

	return len(particles)
}
