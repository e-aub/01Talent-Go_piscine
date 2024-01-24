package piscine

func Unmatch(a []int) int {
	var res int
	for i := 0; i < len(a); i++ {
		for j := i; j < len(a); j++ {
			if a[i] > a[j] {
				a[i], a[j] = a[j], a[i]
			} else {
				continue
			}
		}
	}

	for _, nb := range a {
		for _, nb2 := range a {
			if nb == nb2 {
				continue
			} else {
				res = nb
			}
		}
	}
	return res
}
