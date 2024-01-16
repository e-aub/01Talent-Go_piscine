package piscine

func Checker(nb int) bool {
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

func FindNextPrime(nb int) int {
	if Checker(nb) == false {
		for Checker(nb) == false {
			nb++
		}
	} else if Checker(nb) == true {
		return nb
	}
	return nb
}
