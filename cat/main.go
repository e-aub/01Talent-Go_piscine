package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) == 1 {
		ReadStdin()
		return
	}
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

func ReadStdin() {
	for {
		buffer := make([]byte, 1024)
		n, err := os.Stdin.Read(buffer)
		if err != nil {
			os.Exit(0)
		}

		os.Stdout.Write(buffer[:n])
	}
}
