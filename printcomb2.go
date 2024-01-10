package piscine

import "github.com/01-edu/z01"

func printcomb2() {
	for a := '0'; a <= '9'; a++ {
		for b := '1'; b <= '8'; b++ {
			z01.PrintRune(a)
			z01.PrintRune(b)
			if a != '9' || b != '8' {
				z01.PrintRune(',')
				z01.PrintRune(' ')
			}
		}
	}
	z01.PrintRune('\n')
}
