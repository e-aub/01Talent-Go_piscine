package piscine

func DescendAppendRange(max, min int) []int {
	emptSlice := []int{}
	for i := max; i > min; i-- {
		emptSlice = append(emptSlice, i)
	}
	return emptSlice
}
