package main

import "github.com/01-edu/z01"

func PrintNbr(n int) {
	if n == -9223372036854775808 {
		z01.PrintRune('-')
		z01.PrintRune('9')
		PrintNbr(223372036854775808)
	} else if n < 0 {
		z01.PrintRune('-')
		PrintNbr(-n)
	} else if n >= 0 && n < 10 {
		z01.PrintRune(rune(n + 48))
	} else if n >= 10 {
		PrintNbr(n / 10)
		PrintNbr(n % 10)
	}
}

func main() {
	PrintNbr(-9223372036854775808)
	PrintNbr(0)
	PrintNbr(123)
	z01.PrintRune('\n')
}
