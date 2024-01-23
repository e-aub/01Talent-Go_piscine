package piscine

func IsSorted(f func(a, b int) int, a []int) bool {
	ascending := true
	descending := true
	for i := 0; i < len(a)-1; i++ {
		if f(a[i], a[i+1]) == -1 {
			ascending = false
		} else if f(a[i], a[i+1]) == 0 {
			continue
		} else {
			descending = false
		}
	}
	if !ascending && !descending {
		return false
	}
	return true
}
