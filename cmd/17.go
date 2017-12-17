package cmd

import (
	"github.com/spf13/cobra"
	"strconv"
	"strings"
)

var day17cmd = &cobra.Command{
	Use: "day17",
	Run: run17,
}

func init() {
	RootCmd.AddCommand(day17cmd)
}

func insert17a(stride int, count int) int {
	array := make([]int, 1)
	array[0] = 0
	index := 0
	for i := 1; i <= count; i++ {
		index = (index+stride)%len(array) + 1
		left := array[:index]
		mid := []int{i}
		right := array[index:]
		array = append(left, append(mid, right...)...)
	}
	return array[(index+1)%len(array)]
}

func insert17b(stride int, count int) int {
	index := 0
	value := 0
	for i := 1; i <= count; i++ {
		index = (index+stride)%i + 1
		if index == 1 {
			value = i
		}
	}
	return value
}

func parse17(input string) (int, int) {
	parts := strings.Split(input, ",")
	stride, _ := strconv.Atoi(parts[0])
	count, _ := strconv.Atoi(parts[1])
	return stride, count
}

func compute17a(input string) int {
	stride, count := parse17(input)
	return insert17a(stride, count)
}

func compute17b(input string) int {
	stride, count := parse17(input)
	return insert17b(stride, count)
}

func run17(cmd *cobra.Command, args []string) {
	Test(compute17a, "3,9", 5)
	Test(compute17a, "3,2017", 638)
	input := "382,2017"
	PrintResult(input, compute17a(input))
	inputb := "382,50000000"
	PrintResult(inputb, compute17b(inputb))
}
