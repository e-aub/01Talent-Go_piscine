package piscine

func ShoppingSummaryCounter(str string) map[string]int {
	recipt := make(map[string]int)
	var s string
	var items []string
	if str == "" {
		items = append(items, "\"\"")
	}
	for index, letter := range str {
		if letter == ' ' {
			items = append(items, s)
			s = ""
		} else {
			s = s + string(letter)
		}
		if index == len(str)-1 {
			items = append(items, s)
		}
	}

	for _, item := range items {
		recipt[item] += 1
	}
	return recipt
}
