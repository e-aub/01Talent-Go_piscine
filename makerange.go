package piscine

func MakeRange(min, max int) []int {
	var result []int
	if min == 0 || max == 0 {
		result = []int{0}
		return result
	} else if min > max {
		return result
	}
	result = make([]int, max-min)
	for i := 0; i < max-min; i++ {
		result[i] = min + i
	}
	return result
}
