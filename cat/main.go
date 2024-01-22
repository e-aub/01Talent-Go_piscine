package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args[1:]

	for index, filename := range args {
		Content, err := os.ReadFile(filename)
		if err != nil {
			Printer("ERROR: ")
			errMsg := err.Error()
			Printer(errMsg)
			z01.PrintRune('\n')
			if index == 0 {
				os.Exit(1)
			}
		}
		Printer(string(Content))
	}
}

func Printer(s string) {
	for _, letter := range s {
		z01.PrintRune(letter)
	}
}
