package cmd

import (
	"sort"
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
