package piscine

import (
	"github.com/01-edu/z01"
)

func PrintNbrInOrder(n int) {
	slint := []int{}

	if n == 0 {
		slint = []int{0}
	}
	for n > 0 {
		slint = append(slint, n%10)
		n = n / 10
	}
	for i := 0; i < len(slint); i++ {
		for j := 0; j < len(slint)-1; j++ {
			if slint[j] > slint[j+1] {
				slint[j], slint[j+1] = slint[j+1], slint[j]
			}
		}
	}

	for _, digit := range slint {
		z01.PrintRune(rune(digit) + 48)
	}
}
