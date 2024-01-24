package piscine

func ShoppingSummaryCounter(str string) map[string]int {
	recipt := make(map[string]int)
	var s string
	var items []string

	for _, letter := range str {
		if letter != ' ' {
			s = s + string(letter)
		} else {
			items = append(items, s)
			s = ""
		}
	}
	items = append(items, s)

	for _, item := range items {
		recipt[item]++
	}
	return recipt
}
