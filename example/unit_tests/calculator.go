package example

func Sum(a, b int) int {
	sum := a + b
	if sum < a || sum < b {
		return 0
	}
	return sum
}

func Subtraction(a, b int) int {
	return a - b
}
