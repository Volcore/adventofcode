package cmd

import (
	"bufio"
	"fmt"
	"github.com/spf13/cobra"
	"io/ioutil"
	"strings"
)

var day04cmd = &cobra.Command{
	Use: "day04",
	Run: run04,
}

func init() {
	RootCmd.AddCommand(day04cmd)
}

func checkPassPhrase04(phrase string, sort bool) bool {
	// split into parts
	parts := strings.Split(phrase, " ")
	// put parts in map and check for dup
	m := make(map[string]bool)
	for _, part := range parts {
		if sort {
			part = SortString(part)
		}
		if m[part] == true {
			return false
		}
		m[part] = true
	}
	return true
}

func compute04a(input string) int {
	output := doCompute04a(input)
	str := input
	if len(input) > 200 {
		str = input[0:200] + "..."
	}
	fmt.Println(str, "=>", output)
	return output
}

func doCompute04a(input string) int {
	scanner := bufio.NewScanner(strings.NewReader(input))
	count := 0
	for scanner.Scan() {
		phrase := scanner.Text()
		if checkPassPhrase04(phrase, false) {
			count += 1
		}
	}
	return count
}

func test04a(input string, output int) {
	val := compute04a(input)
	if val != output {
		fmt.Println("Test failed, value should be", output, "but is", val)
	}
}

func compute04b(input string) int {
	output := doCompute04b(input)
	str := input
	if len(input) > 200 {
		str = input[0:200] + "..."
	}
	fmt.Println(str, "=>", output)
	return output
}

func doCompute04b(input string) int {
	scanner := bufio.NewScanner(strings.NewReader(input))
	count := 0
	for scanner.Scan() {
		phrase := scanner.Text()
		if checkPassPhrase04(phrase, true) {
			count += 1
		}
	}
	return count
}

func test04b(input string, output int) {
	val := compute04b(input)
	if val != output {
		fmt.Println("Test failed, value should be", output, "but is", val)
	}
}

func run04(cmd *cobra.Command, args []string) {
	b, err := ioutil.ReadFile("data/04-input.txt")
	if err != nil {
		panic("Input not found!")
	}
	input := strings.Trim(string(b), "\n\r \t")
	test04a("aa bb cc dd ee", 1)
	test04a("aa bb cc dd aa", 0)
	test04a("aa bb cc dd aaa", 1)
	test04a("aa bb cc dd aaa\naa bb cc dd ee", 2)
	test04a("aa bb cc dd aaa\naa bb cc dd aa", 1)
	compute04a(input)
	test04b("abcde fghij", 1)
	test04b("abcde xyz ecdab", 0)
	test04b("a ab abc abd abf abj", 1)
	test04b("iiii oiii ooii oooi oooo", 1)
	test04b("oiii ioii iioi iiio", 0)
	test04b("a ab abc abd abf abj\niiii oiii ooii oooi oooo", 2)
	compute04b(input)
}
