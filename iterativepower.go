package piscine

func IterativePower(nb, power int) int {
	res := 1
	if nb < 0 || power < 0 {
		return 0
	}
	for i := 1; i <= power; i++ {
		res = res * nb
	}
	return res
}
