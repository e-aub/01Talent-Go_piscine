package main

import (
	"github.com/01-edu/z01"
)

type points struct {
	x int
	y int
}

func setPoint(ptr *points) {
	ptr.x = 42
	ptr.y = 21
}

func main() {
	x := points{}
	y := points{}
	setPoint(&x)
	setPoint(&y)
	xEqual := "x = "
	yEqual := " ,y = "
	Print(xEqual)
	PrintNbr(x.x)
	Print(yEqual)
	PrintNbr(y.y)
	z01.PrintRune('\n')
}

func PrintNbr(n int) {
	if n >= 0 && n < 10 {
		z01.PrintRune(rune('0' + n))
	} else if n >= 10 {
		PrintNbr(n / 10)
		PrintNbr(n % 10)
	}
}

func Print(s string) {
	for _, letter := range s {
		z01.PrintRune(letter)
	}
}
