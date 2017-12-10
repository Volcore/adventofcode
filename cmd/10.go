package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"strings"
)

var day10cmd = &cobra.Command{
	Use: "day10",
	Run: run10,
}

func init() {
	RootCmd.AddCommand(day10cmd)
}

func reverseSequence10(memory []int, position int, length int) []int {
	size := len(memory)
	for i := 0; i < length/2; i++ {
		left := (position + i) % size
		right := (position + length - i - 1) % size
		tmp := memory[left]
		memory[left] = memory[right]
		memory[right] = tmp
	}
	return memory
}

func compute10a(input string, size int) int {
	lengths := ParseIntArray(strings.Replace(input, ",", " ", -1))
	memory := make([]int, size)
	for i := 0; i < size; i++ {
		memory[i] = i
	}
	position := 0
	// Permutate
	for idx, length := range lengths {
		skipSize := idx
		// Reverse the sequence length starting at current position
		memory = reverseSequence10(memory, position, length)
		// Move position forward by length
		position = (position + length + skipSize) % size
	}
	// Compute result
	return memory[0] * memory[1]
}

func compute10a5(input string) int {
	return compute10a(input, 5)
}

func compute10a256(input string) int {
	return compute10a(input, 256)
}

type hashContext10 struct {
	position int
	skip     int
}

func (ctx *hashContext10) round10b(box []int, lengths []int) []int {
	for _, length := range lengths {
		// Reverse the sequence length starting at current position
		box = reverseSequence10(box, ctx.position, length)
		// Move position forward by length
		ctx.position = (ctx.position + length + ctx.skip) % 256
		ctx.skip += 1
	}
	return box
}

func compute10b(input string) string {
	box := make([]int, 256)
	for i := 0; i < 256; i++ {
		box[i] = i
	}
	// Parse input
	lengths := make([]int, 0)
	for _, char := range input {
		lengths = append(lengths, int(char))
	}
	lengths = append(lengths, []int{17, 31, 73, 47, 23}...)
	// Run the rounds
	ctx := new(hashContext10)
	for i := 0; i < 64; i++ {
		box = ctx.round10b(box, lengths)
	}
	// Generate output
	digest := ""
	for i := 0; i < 16; i++ {
		// Compute value
		value := uint8(0)
		for j := 0; j < 16; j++ {
			value ^= uint8(box[i*16+j])
		}
		digest += fmt.Sprintf("%02x", value)
	}
	return digest
}

func run10(cmd *cobra.Command, args []string) {
	input := LoadData("data/10-input.txt")
	Test(compute10a5, "3,4,1,5", 12)
	PrintResult(input, compute10a256(input))
	TestS(compute10b, "", "a2582a3a0e66e6e86e3812dcb672a272")
	TestS(compute10b, "AoC 2017", "33efeb34ea91902bb2f59c9920caa6cd")
	TestS(compute10b, "1,2,3", "3efbe78a8d82f29979031a4aa0b16a9d")
	TestS(compute10b, "1,2,4", "63960835bcdc130f0b66d7ff4f6a5a8e")
	PrintResultS(input, compute10b(input))
}
