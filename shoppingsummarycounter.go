package piscine

func ShoppingSummaryCounter(str string) map[string]int {
	recipt := make(map[string]int)
	var s string
	var items []string
	for _, letter := range str {
		if letter == ' ' {
			items = append(items, s)
			s = ""
		} else {
			s = s + string(letter)
		}
	}

	for _, item := range items {
		recipt[item] += 1
	}
	return recipt
}
