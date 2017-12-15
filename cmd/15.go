package cmd

import (
	"github.com/spf13/cobra"
)

var day15cmd = &cobra.Command{
	Use: "day15",
	Run: run15,
}

func init() {
	RootCmd.AddCommand(day15cmd)
}

func rng15(factor uint64, val uint32) uint32 {
	return uint32((uint64(val) * factor) % 2147483647)
}

func compute15a(input string) int {
	initial := ParseIntArray(input)
	state := make([]uint32, 2)
	state[0] = uint32(initial[0])
	state[1] = uint32(initial[1])
	count := 0
	for i := 0; i < 40000000; i++ {
		state[0] = rng15(16807, state[0])
		state[1] = rng15(48271, state[1])
		if state[0]&0xffff == state[1]&0xffff {
			count++
		}
	}
	return count
}

func compute15b(input string) int {
	initial := ParseIntArray(input)
	state := make([]uint32, 2)
	state[0] = uint32(initial[0])
	state[1] = uint32(initial[1])
	count := 0
	for i := 0; i < 5000000; i++ {
		for {
			state[0] = rng15(16807, state[0])
			if state[0]%4 == 0 {
				break
			}
		}
		for {
			state[1] = rng15(48271, state[1])
			if state[1]%8 == 0 {
				break
			}
		}
		if state[0]&0xffff == state[1]&0xffff {
			count++
		}
	}
	return count
}

func run15(cmd *cobra.Command, args []string) {
	input := "722 354"
	Test(compute15a, "65 8921", 588)
	PrintResult(input, compute15a(input))
	Test(compute15b, "65 8921", 309)
	PrintResult(input, compute15b(input))
}
