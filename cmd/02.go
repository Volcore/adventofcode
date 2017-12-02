package cmd

import (
	"bufio"
	"fmt"
	"github.com/spf13/cobra"
	"io/ioutil"
	"strconv"
	"strings"
)

var day02cmd = &cobra.Command{
	Use: "day02",
	Run: run02,
}

func init() {
	RootCmd.AddCommand(day02cmd)
}

func compute02(input string) int {
	fmt.Println("Input is:\n", input)
	count := 0
	scanner := bufio.NewScanner(strings.NewReader(input))
	for scanner.Scan() {
		str := strings.Replace(scanner.Text(), "\t", " ", -1)
		parts := strings.Split(str, " ")
		first, _ := strconv.Atoi(parts[0])
		min := first
		max := first
		for _, part := range parts[1:] {
			val, _ := strconv.Atoi(part)
			if val < min {
				min = val
			}
			if val > max {
				max = val
			}
		}
		diff := max - min
		count += diff
	}
	fmt.Println("Count is: ", count)
	return count
}

func test02(input string, output int) {
	val := compute02(input)
	if val != output {
		fmt.Println("Test failed, value should be", output, "but is", val)
	}
}

func getMultiples02b(parts []string) int {
	for i := 0; i < len(parts); i++ {
		a, _ := strconv.Atoi(parts[i])
		for j := i + 1; j < len(parts); j++ {
			b, _ := strconv.Atoi(parts[j])
			if a%b == 0 {
				return a / b
			}
			if b%a == 0 {
				return b / a
			}
		}
	}
	return 0
}

func compute02b(input string) int {
	fmt.Println("Input is:\n", input)
	count := 0
	scanner := bufio.NewScanner(strings.NewReader(input))
	for scanner.Scan() {
		str := strings.Replace(scanner.Text(), "\t", " ", -1)
		parts := strings.Split(str, " ")
		mult := getMultiples02b(parts)
		count += mult
	}
	fmt.Println("Count is: ", count)
	return count
}

func test02b(input string, output int) {
	val := compute02b(input)
	if val != output {
		fmt.Println("Test failed, value should be", output, "but is", val)
	}
}

func run02(cmd *cobra.Command, args []string) {
	b, err := ioutil.ReadFile("data/02-input.txt")
	if err != nil {
		panic("Input not found!")
	}
	input := strings.Trim(string(b), "\n\r \t")
	test02("5 1 9 5", 8)
	test02("7 5 3", 4)
	test02("2 4 6 8", 6)
	compute02(input)
	test02("5 1 9 5\n7 5 3\n2 4 6 8", 18)
	test02b("5 9 2 8", 4)
	test02b("9 4 7 3", 3)
	test02b("3 8 6 5", 2)
	test02b("5 9 2 8\n9 4 7 3\n3 8 6 5", 9)
	compute02b(input)
}
