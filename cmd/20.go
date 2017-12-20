package cmd

import (
	"bufio"
	"fmt"
	"github.com/spf13/cobra"
	"math"
	"regexp"
	"strconv"
	"strings"
)

var day20cmd = &cobra.Command{
	Use: "day20",
	Run: run20,
}

func init() {
	RootCmd.AddCommand(day20cmd)
}

type particle20 struct {
	p [3]float64
	v [3]float64
	a [3]float64
}

func parse20(input string) []*particle20 {
	re := regexp.MustCompile(`p=<\s*(-?\d+),\s*(-?\d+),\s*(-?\d+)>,\s*v=<\s*(-?\d+),\s*(-?\d+),\s*(-?\d+)>,\s*a=<\s*(-?\d+),\s*(-?\d+),\s*(-?\d+)>`)
	particles := make([]*particle20, 0)
	scanner := bufio.NewScanner(strings.NewReader(input))
	for scanner.Scan() {
		line := scanner.Text()
		match := re.FindStringSubmatch(line)
		if len(match) == 0 {
			if line != "" {
				fmt.Println("ignoring line:", line)
			}
			continue
		}
		p := new(particle20)
		for i := 0; i < 3; i++ {
			p.p[i], _ = strconv.ParseFloat(match[1+i], 64)
			p.v[i], _ = strconv.ParseFloat(match[4+i], 64)
			p.a[i], _ = strconv.ParseFloat(match[7+i], 64)
		}
		particles = append(particles, p)
	}
	return particles
}

func simulate20(particles []*particle20) {
	for _, p := range particles {
		for i := 0; i < 3; i++ {
			p.v[i] += p.a[i]
		}
		for i := 0; i < 3; i++ {
			p.p[i] += p.v[i]
		}
	}
}

func compute20a(input string) int {
	particles := parse20(input)
	// Fast method: find lowest acceleration
	minIdx := 0
	minValue := math.MaxFloat64
	minVel := math.MaxFloat64
	for idx, p := range particles {
		acc := math.Abs(p.a[0]) + math.Abs(p.a[1]) + math.Abs(p.a[2])
		vel := math.Abs(p.v[0]) + math.Abs(p.v[1]) + math.Abs(p.v[2])
		if acc < minValue || (acc == minValue && vel < minVel) {
			minValue = acc
			minVel = vel
			minIdx = idx
		}
	}
	// Slow method: verify by simulation
	for i := 0; i < 10000; i++ {
		simulate20(particles)
	}
	minIdx2 := 0
	minValue2 := math.MaxFloat64
	for idx, p := range particles {
		sum := math.Abs(p.p[0]) + math.Abs(p.p[1]) + math.Abs(p.p[2])
		if sum < minValue2 {
			minValue2 = sum
			minIdx2 = idx
		}
	}
	if minIdx2 != minIdx {
		panic("Verification failed!")
	}
	return minIdx
}

func collide20(particles []*particle20) []*particle20 {
	// Build a map of particle positions
	d := make(map[int]bool)
	m := make(map[string]int, 0)
	for idx, p := range particles {
		bits0 := math.Float64bits(p.p[0])
		bits1 := math.Float64bits(p.p[1])
		bits2 := math.Float64bits(p.p[2])
		key := fmt.Sprintf("%x%x%x\n", bits0, bits1, bits2)
		if oidx, exists := m[key]; exists {
			// add both the current and the map index to the delete map
			d[idx] = true
			d[oidx] = true
			continue
		}
		m[key] = idx
	}
	// Finally delete all elements
	np := make([]*particle20, 0)
	for idx, p := range particles {
		if _, deleted := d[idx]; deleted {
			continue
		}
		np = append(np, p)
	}
	return np
}

func compute20b(input string) int {
	particles := parse20(input)
	for i := 0; i < 1000; i++ {
		simulate20(particles)
		// Check for collisions
		particles = collide20(particles)
	}
	return len(particles)
}

func run20(cmd *cobra.Command, args []string) {
	input := LoadDataRaw("data/20-input.txt")
	test := LoadDataRaw("data/20-test.txt")
	test2 := LoadDataRaw("data/20-test2.txt")
	Test(compute20a, test, 0)
	PrintResult(input, compute20a(input))
	Test(compute20b, test2, 1)
	PrintResult(input, compute20b(input))
}
