package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	// args variable
	args := os.Args
	// Checking for upper flag
	i := 1
	if len(args) == 1 {
		return
	} else if args[1] == "--upper" && len(args) == 1 {
		return
	}
	if args[1] == "--upper" {
		i = 2
	}
	// atoi
	var letter rune
	for ; i < len(args); i++ {
		argsint := 0
		for _, let := range args[i] {
			if let >= '0' && let <= '9' {
				argsint = argsint*10 + int(rune(let)-'0')
			}
		}
		if argsint <= 26 && argsint >= 1 {
			if args[1] == "--upper" {
				letter = rune(argsint + 64)
			} else {
				letter = rune(argsint + 96)
			}
			z01.PrintRune(letter)
		} else {
			z01.PrintRune(' ')
		}

	}
	z01.PrintRune('\n')
}
