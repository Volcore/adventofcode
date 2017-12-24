package cmd

import (
	"bufio"
	"fmt"
	"github.com/spf13/cobra"
	"strconv"
	"strings"
)

var day24cmd = &cobra.Command{
	Use: "day24",
	Run: run24,
}

func init() {
	RootCmd.AddCommand(day24cmd)
}

type comp24 struct {
	port []int
}

func (c *comp24) isOfValue(val int) bool {
	return c.port[0] == val || c.port[1] == val
}

func (c *comp24) otherValue(val int) int {
	if c.port[0] == val {
		return c.port[1]
	}
	return c.port[0]
}

func (c *comp24) String() string {
	return fmt.Sprintf("%d/%d", c.port[0], c.port[1])
}

func parse24(input string) []*comp24 {
	comps := []*comp24{}
	scanner := bufio.NewScanner(strings.NewReader(input))
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, "/")
		c := new(comp24)
		c.port = []int{0, 0}
		c.port[0], _ = strconv.Atoi(parts[0])
		c.port[1], _ = strconv.Atoi(parts[1])
		comps = append(comps, c)
	}
	return comps
}

func slice24(comps []*comp24, i int) []*comp24 {
	return append(append([]*comp24{}, comps[:i]...), comps[i+1:]...)
}

func recurse24(comps []*comp24, sum int, next int) int {
	// find all matches
	max := sum
	for idx, comp := range comps {
		if !comp.isOfValue(next) {
			continue
		}
		other := comp.otherValue(next)
		val := recurse24(slice24(comps, idx), sum+next+other, other)
		if val > max {
			max = val
		}
	}
	return max
}

func compute24a(input string) int {
	comps := parse24(input)
	return recurse24(comps, 0, 0)
}

func run24(cmd *cobra.Command, args []string) {
	test := LoadDataRaw("data/24-test.txt")
	Test(compute24a, test, 31)
	input := LoadDataRaw("data/24-input.txt")
	PrintResult(input, compute24a(input))
}
