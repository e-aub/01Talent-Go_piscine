package piscine

import "math"

func Prime(nb int) bool {
	var res bool
	if nb < 2 {
		return false
	} else if nb == 2 || nb == 3 {
		return true
	}
	for i := 2; i <= int(math.Sqrt(float64(nb))); i++ {
		if nb%i == 0 {
			res = false
		} else {
			res = true
		}
	}
	return res
}

func FindNextPrime(nb int) int {
	for !Prime(nb) {
		nb++
	}
	return nb
}
