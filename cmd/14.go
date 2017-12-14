package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var day14cmd = &cobra.Command{
	Use: "day14",
	Run: run14,
}

func init() {
	RootCmd.AddCommand(day14cmd)
}

func compute14a(input string) int {
	count := 0
	for i := 0; i < 128; i++ {
		rowstr := fmt.Sprintf("%s-%d", input, i)
		hash := compute10b(rowstr)
		for _, c := range hash {
			val := HexToInt(int(c))
			for bit2 := uint8(0); bit2 < 4; bit2++ {
				bit := 3 - bit2
				if val&(1<<bit) != 0 {
					count++
				}
			}
		}
	}
	return count
}

func floodfill14(grid []int, x int, y int, id int) []int {
	if x < 0 || x > 127 || y < 0 || y > 127 {
		return grid
	}
	if grid[x+y*128] != -1 {
		return grid
	}
	grid[x+y*128] = id
	grid = floodfill14(grid, x-1, y, id)
	grid = floodfill14(grid, x, y-1, id)
	grid = floodfill14(grid, x+1, y, id)
	grid = floodfill14(grid, x, y+1, id)
	return grid
}

func compute14b(input string) int {
	// Fill all pieces
	grid := make([]int, 128*128)
	for i := 0; i < 128; i++ {
		rowstr := fmt.Sprintf("%s-%d", input, i)
		hash := compute10b(rowstr)
		j := 0
		for _, c := range hash {
			val := HexToInt(int(c))
			for bit2 := uint8(0); bit2 < 4; bit2++ {
				bit := 3 - bit2
				if val&(1<<bit) != 0 {
					index := j + i*128
					grid[index] = -1
				}
				j++
			}
		}
	}
	// Count regions
	count := 0
	for x := 0; x < 128; x++ {
		for y := 0; y < 128; y++ {
			if grid[x+y*128] == -1 {
				count += 1
				grid = floodfill14(grid, x, y, count)
			}
		}
	}
	return count
}

func run14(cmd *cobra.Command, args []string) {
	Test(compute14a, "flqrgnkx", 8108)
	PrintResult("uugsqrei", compute14a("uugsqrei"))
	Test(compute14b, "flqrgnkx", 1242)
	PrintResult("uugsqrei", compute14b("uugsqrei"))
}
