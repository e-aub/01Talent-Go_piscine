package piscine

func FindNextPrime(nb int) int {
	for !Prime(nb) {
		nb++
	}
	return nb
}

func Prime(nb int) bool {
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
