package chapter3

func factorial(num int) int {
	if num == 1 {
		return 1
	}
	return num * factorial(num-1)
}
