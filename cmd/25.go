package cmd

import (
	"bufio"
	"fmt"
	"github.com/spf13/cobra"
	"regexp"
	"strconv"
	"strings"
)

var day25cmd = &cobra.Command{
	Use: "day25",
	Run: run25,
}

func init() {
	RootCmd.AddCommand(day25cmd)
}

type trans25 struct {
	write int
	dir   int // -1 left +1 right
	next  string
}

type state25 struct {
	trans [2]trans25
}

func (state *state25) String() string {
	return fmt.Sprintf("%+v/%+v", state.trans[0], state.trans[1])
}

type machine25 struct {
	states    map[string]*state25
	cur       string
	stepLimit int
	pos       int
	memory    map[int]int
}

func (machine *machine25) run() {
	for i := 0; i < machine.stepLimit; i++ {
		state := machine.states[machine.cur]
		val := machine.memory[machine.pos]
		// write the new value
		machine.memory[machine.pos] = state.trans[val].write
		// move
		machine.pos += state.trans[val].dir
		// transition
		machine.cur = state.trans[val].next
	}
}

func (machine *machine25) count() int {
	count := 0
	for _, v := range machine.memory {
		count += v
	}
	return count
}

func parse25(input string) machine25 {
	// Regexp that we'll eedj
	emptyRe := regexp.MustCompile(`^\s*$`)
	beginRe := regexp.MustCompile(`^Begin in state (\w+).$`)
	performRe := regexp.MustCompile(`^Perform a diagnostic checksum after (\d+) steps.$`)
	stateRe := regexp.MustCompile(`^In state (\w+):$`)
	currentRe := regexp.MustCompile(`^\s*If the current value is (\d+):$`)
	writeRe := regexp.MustCompile(`^\s*- Write the value (\d+).$`)
	dirRe := regexp.MustCompile(`^\s*- Move one slot to the (\w+).$`)
	continueRe := regexp.MustCompile(`^\s*- Continue with state (\w+).$`)
	// Output tracking
	machine := machine25{}
	machine.states = make(map[string]*state25)
	machine.memory = make(map[int]int)
	var state *state25
	currentValue := 0
	// Start parsing
	scanner := bufio.NewScanner(strings.NewReader(input))
	for scanner.Scan() {
		line := scanner.Text()
		match := beginRe.FindStringSubmatch(line)
		if len(match) != 0 {
			machine.cur = match[1]
			continue
		}
		match = performRe.FindStringSubmatch(line)
		if len(match) != 0 {
			machine.stepLimit, _ = strconv.Atoi(match[1])
			continue
		}
		match = emptyRe.FindStringSubmatch(line)
		if len(match) != 0 {
			continue
		}
		match = stateRe.FindStringSubmatch(line)
		if len(match) != 0 {
			state = new(state25)
			machine.states[match[1]] = state
			continue
		}
		match = currentRe.FindStringSubmatch(line)
		if len(match) != 0 {
			currentValue, _ = strconv.Atoi(match[1])
			continue
		}
		match = writeRe.FindStringSubmatch(line)
		if len(match) != 0 {
			state.trans[currentValue].write, _ = strconv.Atoi(match[1])
			continue
		}
		match = dirRe.FindStringSubmatch(line)
		if len(match) != 0 {
			if match[1] == "left" {
				state.trans[currentValue].dir = -1
			} else {
				state.trans[currentValue].dir = 1
			}
			continue
		}
		match = continueRe.FindStringSubmatch(line)
		if len(match) != 0 {
			state.trans[currentValue].next = match[1]
			continue
		}
		fmt.Println("***", line)
	}
	return machine
}

func compute25a(input string) int {
	machine := parse25(input)
	machine.run()
	return machine.count()
}

func run25(cmd *cobra.Command, args []string) {
	test := LoadDataRaw("data/25-test.txt")
	Test(compute25a, test, 3)
	input := LoadDataRaw("data/25-input.txt")
	PrintResult(input, compute25a(input))
}
