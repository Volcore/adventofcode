package cmd

import (
	"bufio"
	"github.com/spf13/cobra"
	"strconv"
	"strings"
)

var day18cmd = &cobra.Command{
	Use: "day18",
	Run: run18,
}

func init() {
	RootCmd.AddCommand(day18cmd)
}

type vm18 struct {
	vmindex   int
	registers map[string]int
	lastSound int
	sendCount int
	pc        int
	queue     []int
}

type instruction18 struct {
	op  string
	val []string
}

func (inst *instruction18) value(index int, vm *vm18) int {
	val, err := strconv.Atoi(inst.val[index])
	if err == nil {
		return val
	}
	return vm.registers[inst.val[index]]
}

func parse18(input string) []*instruction18 {
	instructions := make([]*instruction18, 0)
	scanner := bufio.NewScanner(strings.NewReader(input))
	for scanner.Scan() {
		row := strings.Split(scanner.Text(), " ")
		instructions = append(instructions, &instruction18{
			op:  row[0],
			val: row[1:],
		})
	}
	return instructions
}

func (vm *vm18) runa(instructions []*instruction18) int {
	inst := instructions[vm.pc]
	switch inst.op {
	case "set":
		vm.registers[inst.val[0]] = inst.value(1, vm)
		break
	case "add":
		vm.registers[inst.val[0]] += inst.value(1, vm)
		break
	case "mul":
		vm.registers[inst.val[0]] *= inst.value(1, vm)
		break
	case "mod":
		vm.registers[inst.val[0]] %= inst.value(1, vm)
		break
	case "snd":
		vm.lastSound = inst.value(0, vm)
		break
	case "rcv":
		if vm.registers[inst.val[0]] != 0 {
			return vm.lastSound
		}
		break
	case "jgz":
		if inst.value(0, vm) > 0 {
			vm.pc += inst.value(1, vm)
			return 0
		}
		break
	default:
		panic("Unknown instruction " + inst.op)
		break
	}
	vm.pc += 1
	return 0
}

func (vm *vm18) runb(instructions []*instruction18, ovm *vm18) bool {
	inst := instructions[vm.pc]
	switch inst.op {
	case "set":
		vm.registers[inst.val[0]] = inst.value(1, vm)
		break
	case "add":
		vm.registers[inst.val[0]] += inst.value(1, vm)
		break
	case "mul":
		vm.registers[inst.val[0]] *= inst.value(1, vm)
		break
	case "mod":
		vm.registers[inst.val[0]] %= inst.value(1, vm)
		break
	case "snd":
		// Try parsing as value, otherwise get register
		ovm.queue = append(ovm.queue, inst.value(0, vm))
		vm.sendCount += 1
		break
	case "rcv":
		// Can't receive, queue empty.
		if len(vm.queue) == 0 {
			return false
		}
		vm.registers[inst.val[0]] = vm.queue[0]
		vm.queue = vm.queue[1:]
		break
	case "jgz":
		if inst.value(0, vm) > 0 {
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

func (vm *vm18) blocked() bool {
	return len(vm.queue) == 0
}

func compute18a(input string) int {
	instructions := parse18(input)
	vm := new(vm18)
	vm.registers = make(map[string]int, 0)
	for {
		val := vm.runa(instructions)
		if val != 0 {
			return val
		}
	}
	return 0
}

func compute18b(input string) int {
	instructions := parse18(input)
	vms := make([]*vm18, 2)
	for i := 0; i < 2; i++ {
		vms[i] = new(vm18)
		vms[i].registers = make(map[string]int, 0)
		vms[i].registers["p"] = i
		vms[i].vmindex = i
	}
	// simulate until deadlock
	for {
		// simulate each vm until it is blocked
		for i := 0; i < 2; i++ {
			for vms[i].runb(instructions, vms[(i+1)%2]) {
			}
		}
		// Both vms have empty queues? abort!
		if vms[0].blocked() && vms[1].blocked() {
			break
		}
	}
	return vms[1].sendCount
}

func run18(cmd *cobra.Command, args []string) {
	input := LoadData("data/18-input.txt")
	test := LoadData("data/18-test.txt")
	test2 := LoadData("data/18-test2.txt")
	Test(compute18a, test, 4)
	PrintResult(input, compute18a(input))
	Test(compute18b, test2, 3)
	PrintResult(input, compute18b(input))
}
