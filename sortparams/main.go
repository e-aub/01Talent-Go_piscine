package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args
	for i := 1; i < len(args); i++ {
		for j := 1; j < len(args)-1; j++ {
			if args[i] < args[j] {
				args[i], args[j] = args[j], args[i]
			}
		}
	}
	for index := 1; index < len(args); index++ {
		for _, let := range args[index] {
			z01.PrintRune(let)
			z01.PrintRune('\n')
		}
	}
}
