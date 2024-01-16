package piscine

func FindNextPrime(nb int) int {
	for !IsPrime(nb) {
		nb++
	}
	return nb
}

func IsPrime(nb int) bool {
	if nb < 2 {
		return false
	}
	for i := 2; i < nb; i++ {
		if nb%i == 0 {
			return false
		}
	}
	return true
}
