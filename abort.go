package piscine

func Abort(a, b, c, d, e int) int {
	sl := []int{a, b, c, d, e}
	for i := 0; i < len(sl); i++ {
		for j := i; j < len(sl); j++ {
			if sl[i] > sl[j] {
				sl[i], sl[j] = sl[j], sl[i]
			}
		}
	}
	return int(sl[2])
}
