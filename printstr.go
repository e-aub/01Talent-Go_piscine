package piscine

import "github.com/01-edu/z01"

func PrintStr(s string) {
	for _, char := range s {
		z01.PrintRune(char)
	}
}

func main() {
	PrintStr("Hello World!")
}
