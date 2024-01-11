package piscine

import "github.com/01-edu/z01"

func PrintNbr(n int) {
	MinVal := -9223372036854775808
	if n == MinVal {
		PrintNbr(-922337203685477580)
		PrintNbr(8)
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
