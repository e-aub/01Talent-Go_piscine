package piscine

func IterativeFactorial(nb int) int {
	fact := 1
	for i := 1; i <= nb; i++ {
		fact = fact * i
	}
	return fact
}
