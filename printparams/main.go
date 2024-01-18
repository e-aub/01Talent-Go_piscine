package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args
	for i := 1; i <= len(args)-1; i++ {
		for _, let := range args[i] {
			z01.PrintRune(let)
		}
		z01.PrintRune('\n')
	}
}
