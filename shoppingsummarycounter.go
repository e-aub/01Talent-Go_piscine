package piscine

func ShoppingSummaryCounter(str string) map[string]int {
	Burger := 0
	Water := 0
	Carrot := 0
	Coffee := 0
	Chips := 0
	var recipt map[string]int = map[string]int{
		"Burger": 0,
		"Water":  0,
		"Carrot": 0,
		"Coffee": 0,
		"Chips":  0,
	}

	var s string
	var items []string
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
		if item == "Burger" {
			Burger++
			recipt["Burger"] = Burger
		} else if item == "Water" {
			Water++
			recipt["Water"] = Water
		} else if item == "Carrot" {
			Carrot++
			recipt["Carrot"] = Carrot
		} else if item == "Coffee" {
			Coffee++
			recipt["Coffee"] = Coffee
		} else if item == "Chips" {
			Chips++
			recipt["Chips"] = Chips
		}
	}
	return recipt
}
