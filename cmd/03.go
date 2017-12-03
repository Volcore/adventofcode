package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"math"
)

var day03cmd = &cobra.Command{
	Use: "day03",
	Run: run03,
}

func init() {
	RootCmd.AddCommand(day03cmd)
}

func compute03(input int) int {
	output := doCompute03(input)
	fmt.Println(input, "=>", output)
	return output
}
func doCompute03(input int) int {
	if input == 1 {
		return 0
	}
	// Compute the current level
	level := int(math.Ceil(0.5*math.Sqrt(float64(input)) + 0.5))
	// Level is the absolute of the first manhattance distance coordinate
	// How long is one side of the level
	sideLength := 2*level - 1
	// The start index of the level
	start := SqrInt(2*(level-1) - 1)
	// Compute the offset in a circle
	circleOffset := input - start - 1
	// compute the offset on a specific side
	sideOffset := circleOffset % (sideLength - 1)
	// we're almost at the second parameter for the manhattan distance
	// all we need to do is adjust the offset for that it starts off-center
	offset := sideOffset - (sideLength/2 - 1)
	// Compute the manhattan distance
	return (level - 1) + AbsInt(offset)
}

func test03(input int, output int) {
	val := compute03(input)
	if val != output {
		fmt.Println("Test failed, value should be", output, "but is", val)
	}
}

func makeIndex03b(x int, y int) int {
	return x + (y << 16)
}

func compute03b(target int) int {
	output := doCompute03b(target)
	fmt.Println(target, "=>", output)
	return output
}

func nextMapStep03b(m map[int]int, i int) int {
	// Compute coordinates
	level := int(math.Ceil(0.5*math.Sqrt(float64(i+1)) + 0.5))
	sideLength := 2*level - 1
	start := SqrInt(2*(level-1) - 1)
	circleOffset := i - start
	sideOffset := circleOffset % (sideLength - 1)
	offset := sideOffset - (sideLength/2 - 1)
	sideIndex := circleOffset / (sideLength - 1)
	x := 0
	y := 0
	switch sideIndex {
	case 0:
		x = level - 1
		y = offset
	case 1:
		x = -offset
		y = level - 1
	case 2:
		x = -(level - 1)
		y = -offset
	case 3:
		x = offset
		y = -(level - 1)
	}
	// Compute location
	index := makeIndex03b(x, y)
	// Fetch 8 patch around index
	sum := 0
	for dx := -1; dx <= 1; dx++ {
		for dy := -1; dy <= 1; dy++ {
			idx := makeIndex03b(x+dx, y+dy)
			sum += m[idx]
		}
	}
	m[index] = sum
	return sum
}

func doCompute03b(target int) int {
	m := make(map[int]int)
	m[0] = 1
	sum := 1
	for i := 1; i <= target; i++ {
		sum = nextMapStep03b(m, i)
	}
	// Return the last sum
	return sum
}

func find03b(target int) int {
	output := doFind03b(target)
	fmt.Println(target, "found", output)
	return output
}

func findTest03b(target int, ref int) int {
	output := doFind03b(target)
	fmt.Println(target, ref, output)
	if output != ref {
		fmt.Println("Test failed, value should be", ref, "but is", output)
	}
	return output
}

func doFind03b(value int) int {
	m := make(map[int]int)
	m[0] = 1
	i := 1
	for {
		sum := nextMapStep03b(m, i)
		if sum > value {
			return sum
		}
		i += 1
	}
	return 0
}

func test03b(index int, output int) {
	val := compute03b(index)
	if val != output {
		fmt.Println("Test failed, value for index", index, "should be", output, "but is", val)
	}
}

func run03(cmd *cobra.Command, args []string) {
	test03(1, 0)
	test03(2, 1)
	test03(3, 2)
	test03(4, 1)
	test03(5, 2)
	test03(6, 1)
	test03(7, 2)
	test03(8, 1)
	test03(9, 2)
	test03(12, 3)
	test03(17, 4)
	test03(21, 4)
	test03(22, 3)
	test03(23, 2)
	test03(1024, 31)
	compute03(325489)
	test03b(0, 1)
	test03b(1, 1)
	test03b(2, 2)
	test03b(3, 4)
	test03b(4, 5)
	test03b(5, 10)
	test03b(6, 11)
	test03b(7, 23)
	test03b(8, 25)
	test03b(9, 26)
	test03b(10, 54)
	test03b(11, 57)
	test03b(12, 59)
	findTest03b(24, 25)
	findTest03b(56, 57)
	findTest03b(57, 59)
	findTest03b(58, 59)
	find03b(325489)
}
