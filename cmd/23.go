package cmd

import (
	"bufio"
	"fmt"
	"github.com/spf13/cobra"
	"strconv"
	"strings"
)

var day23cmd = &cobra.Command{
	Use: "day23",
	Run: run23,
}

func init() {
	RootCmd.AddCommand(day23cmd)
}

type vm23 struct {
	registers map[string]int
	pc        int
	mulcount  int
}

type instruction23 struct {
	op  string
	val []string
}

func (inst *instruction23) value(index int, vm *vm23) int {
	val, err := strconv.Atoi(inst.val[index])
	if err == nil {
		return val
	}
	return vm.registers[inst.val[index]]
}

func parse23(input string) []*instruction23 {
	instructions := make([]*instruction23, 0)
	scanner := bufio.NewScanner(strings.NewReader(input))
	for scanner.Scan() {
		row := strings.Split(scanner.Text(), " ")
		instructions = append(instructions, &instruction23{
			op:  row[0],
			val: row[1:],
		})
	}
	return instructions
}

func (vm *vm23) print() {
	fmt.Print(" a=", vm.registers["a"])
	fmt.Print(" b=", vm.registers["b"])
	fmt.Print(" c=", vm.registers["c"])
	fmt.Print(" d=", vm.registers["d"])
	fmt.Print(" e=", vm.registers["e"])
	fmt.Print(" f=", vm.registers["f"])
	fmt.Print(" g=", vm.registers["g"])
	fmt.Print(" h=", vm.registers["h"])
	fmt.Println("")
}

func (vm *vm23) run(instructions []*instruction23) bool {
	if vm.pc < 0 || vm.pc >= len(instructions) {
		return false
	}
	inst := instructions[vm.pc]
	switch inst.op {
	case "set":
		vm.registers[inst.val[0]] = inst.value(1, vm)
		break
	case "sub":
		vm.registers[inst.val[0]] -= inst.value(1, vm)
		break
	case "mul":
		vm.registers[inst.val[0]] *= inst.value(1, vm)
		vm.mulcount++
		break
	case "jnz":
		if inst.value(0, vm) != 0 {
			vm.pc += inst.value(1, vm)
			return true
		}
		break
	default:
		panic("Unknown instruction " + inst.op)
		break
	}
	vm.pc += 1
	return true
}

func compute23a(input string) int {
	instructions := parse23(input)
	vm := new(vm23)
	vm.registers = make(map[string]int, 0)
	for {
		if !vm.run(instructions) {
			return vm.mulcount
		}
	}
	return 0
}

func compute23b(input string) int {
	number := 57
	multiplier := 100
	b_adder := 100000
	c_adder := 17000
	b := number*multiplier + b_adder
	c := b + c_adder
	h := 0
	for b <= c {
		for d := 2; d != b; d++ {
			if b%d == 0 {
				h++
				break
			}
		}
		b += 17
	}
	return h
}

func run23(cmd *cobra.Command, args []string) {
	input := LoadDataRaw("data/23-input.txt")
	PrintResult(input, compute23a(input))
	input2 := LoadDataRaw("data/23-input.txt")
	PrintResult(input, compute23b(input2))
}
