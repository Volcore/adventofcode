package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var day09cmd = &cobra.Command{
	Use: "day09",
	Run: run09,
}

func init() {
	RootCmd.AddCommand(day09cmd)
}

type group09 struct {
	groups []*group09
	cancel int
}

func parse09(input string) *group09 {
	group := new(group09)
	stack := make([]*group09, 0)
	stack = append(stack, group)
	// Always assume first char is {
	idx := 1
	for {
		if idx >= len(input) {
			break
		}
		c := input[idx]
		idx += 1
		switch c {
		case '!':
			idx += 1
			break
		case ',':
			// Ignore commas
			break
		case '{':
			g := new(group09)
			stack[len(stack)-1].groups = append(stack[len(stack)-1].groups, g)
			stack = append(stack, g)
			break
		case '}':
			// Remove last group element
			stack = stack[:len(stack)-1]
			break
		case '<':
			// Continue scan until first close
			for {
				if idx >= len(input) {
					break
				}
				c2 := input[idx]
				idx += 1
				if c2 == '!' {
					idx += 1
					continue
				}
				if c2 == '>' {
					break
				}
				stack[len(stack)-1].cancel += 1
			}
			break
		default:
			fmt.Println(c, input[idx-1:])
			break
		}
	}
	return group
}

func count09a(group *group09, depth int) int {
	count := depth
	for _, sub := range group.groups {
		count += count09a(sub, depth+1)
	}
	return count
}

func compute09a(input string) int {
	group := parse09(input)
	return count09a(group, 1)
}

func count09b(group *group09) int {
	count := group.cancel
	for _, sub := range group.groups {
		count += count09b(sub)
	}
	return count
}

func compute09b(input string) int {
	group := parse09(input)
	return count09b(group)
}

func run09(cmd *cobra.Command, args []string) {
	input := LoadData("data/09-input.txt")
	Test(compute09a, "{}", 1)
	Test(compute09a, "{{}}", 3)
	Test(compute09a, "{{{}}}", 6)
	Test(compute09a, "{{},{}}", 5)
	Test(compute09a, "{{{},{},{{}}}}", 16)
	Test(compute09a, "{<a>,<a>,<a>,<a>}", 1)
	Test(compute09a, "{{<ab>},{<ab>},{<ab>},{<ab>}}", 9)
	Test(compute09a, "{{<!!>},{<!!>},{<!!>},{<!!>}}", 9)
	Test(compute09a, "{{<a!>},{<a!>},{<a!>},{<ab>}}", 3)
	PrintResult(input, compute09a(input))
	Test(compute09b, "{}", 0)
	Test(compute09b, "{<a>}", 1)
	Test(compute09b, "{<abc>}", 3)
	Test(compute09b, "{<a!c>}", 1)
	Test(compute09b, "{<!cab>}", 2)
	Test(compute09b, "{<a>,<a>,<a>,<a>}", 4)
	Test(compute09b, "{{<ab>},{<ab>},{<ab>},{<ab>}}", 8)
	Test(compute09b, "{{<!!>},{<!!>},{<!!>},{<!!>}}", 0)
	Test(compute09b, "{{<a!>},{<a!>},{<a!>},{<ab>}}", 17)
	PrintResult(input, compute09b(input))
}
