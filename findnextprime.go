package main

import "fmt"

func FindNextPrime(nb int) int {
	if IsPrime(nb) == false {
		for IsPrime(nb) == false {
			nb++
		}
	} else if IsPrime(nb) == true {
		return nb
	}
	return nb
}

func IsPrime(nb int) bool {
	var res bool
	if nb < 1 {
		return false
	} else if nb == 2 || nb == 3 {
		return true
	}
	for i := 2; i <= nb/2; i++ {
		if nb%i == 0 {
			res = false
			break
		} else {
			res = true
		}
	}
	return res
}

func main() {
	fmt.Println(FindNextPrime(90))
	fmt.Println(FindNextPrime(62))
}
