package piscine

func Compact(ptr *[]string) int {
	var newSlice []string
	for _, in := range *ptr {
		if in == "" {
			continue
		} else {
			newSlice = append(newSlice, in)
		}
	}
	*ptr = newSlice
	return len(newSlice)
}
