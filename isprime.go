package main

import "math"

func IsPrime(nb int) bool {
	var res bool
	for i := 2; i <= int(math.Sqrt(float64(nb))); i++ {
		if nb%i == 0 {
			res = false
		} else {
			res = true
		}
	}
	return res
}
