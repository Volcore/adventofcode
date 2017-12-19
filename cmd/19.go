package cmd

import (
	"bufio"
	"github.com/spf13/cobra"
	"strings"
)

var day19cmd = &cobra.Command{
	Use: "day19",
	Run: run19,
}

func init() {
	RootCmd.AddCommand(day19cmd)
}

type pos19 struct {
	x     int
	y     int
	steps int
}

func (pos *pos19) applyDir(dir int) {
	switch dir {
	case 0:
		pos.y--
		break
	case 1:
		pos.x++
		break
	case 2:
		pos.y++
		break
	case 3:
		pos.x--
		break
	}
	pos.steps++
}

func (pos *pos19) valid(maze [][]rune) bool {
	return pos.x >= 0 && pos.y >= 0 && pos.y < len(maze) && pos.x < len(maze[0])
}

func (pos *pos19) canWalkInDir(maze [][]rune, dir int) bool {
	// Compute position in that direction and make sure it's okay
	p := *pos
	p.applyDir(dir)
	if !p.valid(maze) {
		return false
	}
	// Check symbol
	r := maze[p.y][p.x]
	if r >= 'A' && r <= 'Z' {
		return true
	}
	if r == '+' {
		return true
	}
	// only can go vertical if up or down
	if r == '|' {
		return dir == 0 || dir == 2
	}
	// only can go horizontal if left or right
	if r == '-' {
		return dir == 1 || dir == 3
	}
	return false
}

func oppositeDir19(dir int) int {
	switch dir {
	case 0:
		return 2
	case 1:
		return 3
	case 2:
		return 0
	case 3:
		return 1
	}
	return 0
}

func parse19(input string) [][]rune {
	maze := make([][]rune, 0)
	scanner := bufio.NewScanner(strings.NewReader(input))
	for scanner.Scan() {
		row := make([]rune, 0)
		for _, r := range scanner.Text() {
			row = append(row, r)
		}
		maze = append(maze, row)
	}
	return maze
}

func findStart19(maze [][]rune) *pos19 {
	pos := new(pos19)
	for idx, r := range maze[0] {
		if r == 124 {
			pos.x = idx
			return pos
		}
	}
	panic("No entrance found!")
	return pos
}

func walk19(maze [][]rune, pos *pos19) string {
	// directions: 0 = up, 1 = right, 2 = down, 3 = left
	direction := 2
	path := ""
	for {
		if !pos.valid(maze) {
			return path
		}
		r := maze[pos.y][pos.x]
		switch r {
		// these are the lines where our only option is to keep going anyways
		case '|':
			break
		case '-':
			break
		case '+':
			for i := 0; i < 4; i++ {
				// don't go back, also don't go same dir again
				if oppositeDir19(direction) == i || direction == i {
					continue
				}
				if pos.canWalkInDir(maze, i) {
					direction = i
					break
				}
			}
			break
		case ' ':
			return path
		default:
			if r >= 'A' && r <= 'Z' {
				path += string(r)
				break
			}
			panic("unknown element, do not know how to proceed")
		}
		// move along direction
		pos.applyDir(direction)
	}
}

func compute19a(input string) string {
	maze := parse19(input)
	// Find startig point
	pos := findStart19(maze)
	// Continue through the maze, gathering the letters
	return walk19(maze, pos)
}

func compute19b(input string) int {
	maze := parse19(input)
	// Find startig point
	pos := findStart19(maze)
	// Continue through the maze
	walk19(maze, pos)
	return pos.steps
}

func run19(cmd *cobra.Command, args []string) {
	input := LoadDataRaw("data/19-input.txt")
	test := LoadDataRaw("data/19-test.txt")
	TestS(compute19a, test, "ABCDEF")
	PrintResultS(input, compute19a(input))
	Test(compute19b, test, 38)
	PrintResult(input, compute19b(input))
}
