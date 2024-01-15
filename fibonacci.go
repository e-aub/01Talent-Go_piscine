package piscine

func Fibonacci(index int) int {
	res := 0
	if index < 0 {
		return -1
	} else if index == 0 {
		return 0
	} else if index <= 2 {
		return 1
	} else if index > 1 {
		res = Fibonacci(index-1) + Fibonacci(index-2)
	}
	return res
}
