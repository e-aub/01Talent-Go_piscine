package piscine

func Unmatch(a []int) int {
	var count int
	for i := 0; i < len(a); i++ {
		for j := 0; j < len(a); j++ {
			if a[i] == a[j] {
				count++
			}
		}
		if count%2 != 0 {
			return a[i]
		}
	}
	return -1
}
