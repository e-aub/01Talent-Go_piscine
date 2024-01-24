package piscine

func StringToIntSlice(str string) []int {
	var intSlice []int
	for _, letter := range str {
		intSlice = append(intSlice, int(byte(letter)))
	}
	return intSlice
}
