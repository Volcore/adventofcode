package cmd

import (
	"bufio"
	"github.com/spf13/cobra"
	"math"
	"strconv"
	"strings"
)

var day08cmd = &cobra.Command{
	Use: "day08",
	Run: run08,
}

func init() {
	RootCmd.AddCommand(day08cmd)
}

type instruction08 struct {
	register  string
	op        string
	value     int
	cmp_a     string
	condition string
	cmp_b     int
}

func parseProgram08(input string) []*instruction08 {
	array := make([]*instruction08, 0)
	scanner := bufio.NewScanner(strings.NewReader(input))
	for scanner.Scan() {
		parts := strings.Split(scanner.Text(), " ")
		inst := new(instruction08)
		inst.register = parts[0]
		inst.op = parts[1]
		inst.value, _ = strconv.Atoi(parts[2])
		inst.cmp_a = parts[4]
		inst.condition = parts[5]
		inst.cmp_b, _ = strconv.Atoi(parts[6])
		array = append(array, inst)
	}
	return array
}

func runProgram08(prog []*instruction08, partTwo bool) int {
	registers := make(map[string]int)
	largest := math.MinInt32
	for _, inst := range prog {
		// Check conditional
		a := registers[inst.cmp_a]
		b := inst.cmp_b
		execute := false
		switch inst.condition {
		case ">":
			execute = (a > b)
			break
		case "<":
			execute = (a < b)
			break
		case ">=":
			execute = (a >= b)
			break
		case "<=":
			execute = (a <= b)
			break
		case "==":
			execute = (a == b)
			break
		case "!=":
			execute = (a != b)
			break
		}
		if !execute {
			continue
		}
		// execute operation
		reg := registers[inst.register]
		if inst.op == "inc" {
			reg += inst.value
		} else {
			reg -= inst.value
		}
		if reg > largest {
			largest = reg
		}
		registers[inst.register] = reg
	}
	if partTwo {
		return largest
	}
	// Check registers for largest one
	largest = math.MinInt32
	for _, val := range registers {
		if val > largest {
			largest = val
		}
	}
	return largest
}

func compute08a(input string) int {
	prog := parseProgram08(input)
	return runProgram08(prog, false)
}

func compute08b(input string) int {
	prog := parseProgram08(input)
	return runProgram08(prog, true)
}

func run08(cmd *cobra.Command, args []string) {
	test := LoadData("data/08-test.txt")
	input := LoadData("data/08-input.txt")
	Test(compute08a, test, 1)
	PrintResult(input, compute08a(input))
	Test(compute08b, test, 10)
	PrintResult(input, compute08b(input))
}
