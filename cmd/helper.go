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

func PrintIntArray(array []int) string {
	return fmt.Sprint(array)
}
