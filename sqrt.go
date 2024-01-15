package piscine

func Sqrt(nb int) int {
	root := 0
	if nb < 0 {
		return root
	}
	for i := 0; i <= nb; i++ {
		if nb == i*i {
			root = i
			break
		}
	}
	return root
}
