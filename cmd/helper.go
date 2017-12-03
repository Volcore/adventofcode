package cmd

func SqrInt(i int) int {
	return i * i
}

func AbsInt(i int) int {
	if i < 0 {
		return -i
	}
	return i
}
