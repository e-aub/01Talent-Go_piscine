package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args[0]
	for i, let := range args {
		if i > 0 {
			z01.PrintRune(rune(let))
		}
	}
	z01.PrintRune('\n')
}
