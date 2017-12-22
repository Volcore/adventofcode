package cmd

import (
	"bufio"
	"fmt"
	"github.com/spf13/cobra"
	"strings"
)

var day22cmd = &cobra.Command{
	Use: "day22",
	Run: run22,
}

func init() {
	RootCmd.AddCommand(day22cmd)
}

const (
	kUp    = iota
	kRight = iota
	kDown  = iota
	kLeft  = iota
)

const (
	kClean    = iota
	kWeak     = iota
	kInfected = iota
	kFlagged  = iota
)

type context22 struct {
	grid     map[string]int
	x        int
	y        int
	dir      int
	infected int
}

func coord22(x int, y int) string {
	return fmt.Sprintf("%d|%d", x, y)
}

func parse22(input string) *context22 {
	ctx := new(context22)
	ctx.grid = make(map[string]int, 0)
	scanner := bufio.NewScanner(strings.NewReader(input))
	y := 0
	for scanner.Scan() {
		x := 0
		line := scanner.Text()
		for _, r := range line {
			if r == '#' {
				ctx.grid[coord22(x, y)] = kInfected
			}
			x++
		}
		y++
	}
	ctx.x = y / 2
	ctx.y = y / 2
	return ctx
}

func (ctx *context22) draw(min int, max int) {
	for y := min; y <= max; y++ {
		for x := min; x <= max; x++ {
			cur := coord22(x, y)
			if ctx.grid[cur] == kInfected {
				if ctx.x == x && ctx.y == y {
					fmt.Print("%")
				} else {
					fmt.Print("#")
				}
			} else {
				if ctx.x == x && ctx.y == y {
					fmt.Print(",")
				} else {
					fmt.Print(".")
				}

			}
		}
		fmt.Println("")
	}
	fmt.Println("")
}

func (ctx *context22) move() {
	switch ctx.dir {
	case kUp:
		ctx.y--
		break
	case kDown:
		ctx.y++
		break
	case kRight:
		ctx.x++
		break
	case kLeft:
		ctx.x--
		break
	}
}

func (ctx *context22) left() {
	ctx.dir = (ctx.dir - 1 + 4) % 4
}

func (ctx *context22) right() {
	ctx.dir = (ctx.dir + 1) % 4
}

func (ctx *context22) reverse() {
	ctx.dir = (ctx.dir + 2) % 4
}

func (ctx *context22) simulatea() {
	// Check if current node is infected
	cur := coord22(ctx.x, ctx.y)
	infected := ctx.grid[cur] == kInfected
	// Turn left or right
	if infected {
		// Turn right
		ctx.right()
		// clean
		ctx.grid[cur] = kClean
	} else {
		// Turn left
		ctx.left()
		// infect
		ctx.grid[cur] = kInfected
		ctx.infected++
	}
	// Move forward
	ctx.move()
}

func (ctx *context22) simulateb() {
	// Check if current node is infected
	cur := coord22(ctx.x, ctx.y)
	// Turn left or right
	switch ctx.grid[cur] {
	case kClean:
		ctx.grid[cur] = kWeak
		ctx.left()
		break
	case kWeak:
		ctx.grid[cur] = kInfected
		ctx.infected++
		break
	case kInfected:
		ctx.grid[cur] = kFlagged
		ctx.right()
		break
	case kFlagged:
		ctx.grid[cur] = kClean
		ctx.reverse()
		break
	}
	// Move forward
	ctx.move()
}

func compute22a(input string) int {
	ctx := parse22(input)
	for i := 0; i < 10000; i++ {
		ctx.simulatea()
	}
	return ctx.infected
}

func compute22b(input string) int {
	ctx := parse22(input)
	for i := 0; i < 10000000; i++ {
		ctx.simulateb()
	}
	return ctx.infected
}

func run22(cmd *cobra.Command, args []string) {
	input := LoadDataRaw("data/22-input.txt")
	Test(compute22a, "..#\n#..\n...\n", 5587)
	PrintResult(input, compute22a(input))
	Test(compute22b, "..#\n#..\n...\n", 2511944)
	PrintResult(input, compute22b(input))
}
