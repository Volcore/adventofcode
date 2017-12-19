package cmd

import (
	"fmt"
	"io/ioutil"
	"sort"
	"strconv"
	"strings"
)

func SqrInt(i int) int {
	return i * i
}

func AbsInt(i int) int {
	if i < 0 {
		return -i
	}
	return i
}

func SortString(w string) string {
	s := strings.Split(w, "")
	sort.Strings(s)
	return strings.Join(s, "")
}

func PrintResult(input string, output int) {
	// Crop input and output
	input = strings.Replace(input, "\n", " ", -1)
	if len(input) > 50 {
		input = input[:50] + "..."
	}
	fmt.Println(input, "=>", output)
}

func PrintResultS(input string, output string) {
	// Crop input and output
	input = strings.Replace(input, "\n", " ", -1)
	if len(input) > 50 {
		input = input[:50] + "..."
	}
	fmt.Println(input, "=>", output)
}

func LoadData(path string) string {
	b, err := ioutil.ReadFile(path)
	if err != nil {
		panic("Input " + path + " not found!")
	}
	return strings.Trim(string(b), "\n\r \t")
}

func LoadDataRaw(path string) string {
	b, err := ioutil.ReadFile(path)
	if err != nil {
		panic("Input " + path + " not found!")
	}
	return string(b)
}

func ParseIntArray(input string) []int {
	sarray := strings.Split(input, " ")
	iarray := make([]int, len(sarray))
	for idx, s := range sarray {
		iarray[idx], _ = strconv.Atoi(s)
	}
	return iarray
}

func Test(f func(string) int, input string, output int) {
	res := f(input)
	PrintResult(input, res)
	if res != output {
		fmt.Println("Test failed, value should be", output, "but is", res)
	}
}

func TestS(f func(string) string, input string, output string) {
	res := f(input)
	PrintResultS(input, res)
	if res != output {
		fmt.Println("Test failed, value should be", output, "but is", res)
	}
}

func HexToInt(c int) int {
	if c >= '0' && c <= '9' {
		return c - '0'
	}
	if c >= 'a' && c <= 'f' {
		return c - 'a' + 10
	}
	return 0
}

func PrintIntArray(array []int) string {
	return fmt.Sprint(array)
}
