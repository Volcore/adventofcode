package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"strings"
)

var day11cmd = &cobra.Command{
	Use: "day11",
	Run: run11,
}

func init() {
	RootCmd.AddCommand(day11cmd)
}

type coord11 struct {
	x int
	y int
	z int
}

func (pos *coord11) distance() int {
	return (AbsInt(pos.x) + AbsInt(pos.y) + AbsInt(pos.z)) / 2
}

func compute11(input string, max bool) int {
	dirs := strings.Split(input, ",")
	pos := coord11{}
	farthest := 0
	for _, dir := range dirs {
		switch dir {
		case "n":
			pos.y += 1
			pos.z -= 1
			break
		case "s":
			pos.y -= 1
			pos.z += 1
			break
		case "ne":
			pos.x += 1
			pos.z -= 1
			break
		case "sw":
			pos.x -= 1
			pos.z += 1
			break
		case "nw":
			pos.x -= 1
			pos.y += 1
			break
		case "se":
			pos.x += 1
			pos.y -= 1
			break
		default:
			fmt.Println("unknown direction", dir)
			break
		}
		distance := pos.distance()
		if distance > farthest {
			farthest = distance
		}
	}
	// Compute distance
	if max {
		return farthest
	} else {
		return pos.distance()
	}
}

func compute11a(input string) int {
	return compute11(input, false)
}

func compute11b(input string) int {
	return compute11(input, true)
}

func run11(cmd *cobra.Command, args []string) {
	input := LoadData("data/11-input.txt")
	Test(compute11a, "ne,ne,ne", 3)
	Test(compute11a, "ne,ne,sw,sw", 0)
	Test(compute11a, "se,se", 2)
	Test(compute11a, "ne,ne,s,s", 2)
	Test(compute11a, "s,s,sw", 3)
	Test(compute11a, "se,sw,se,sw,sw", 3)
	PrintResult(input, compute11a(input))
	Test(compute11b, "ne,ne,ne", 3)
	Test(compute11b, "ne,ne,sw,sw", 2)
	Test(compute11b, "se,se", 2)
	Test(compute11b, "ne,ne,s,s", 2)
	Test(compute11b, "s,s,sw", 3)
	Test(compute11b, "se,sw,se,sw,sw", 3)
	PrintResult(input, compute11b(input))
}
