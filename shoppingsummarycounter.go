package piscine

func ShoppingSummaryCounter(str string) map[string]int {
	Burger := 0
	Water := 0
	Carrot := 0
	Coffee := 0
	Chips := 0
	e := 0
	space := 0
	var recipt map[string]int = map[string]int{
		"Burger": 0,
		"Water":  0,
		"Carrot": 0,
		"Coffee": 0,
		"Chips":  0,
		"e":      0,
		"\"\"":   0,
	}

	var s string
	var items []string
	for index, letter := range str {
		if letter == ' ' {
			items = append(items, s)
			s = ""
		} else if letter == ' ' && str[index-1] == ' ' {
			items = append(items, "")
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
		} else if item == "" {
			space++
			recipt["\"\""] = space
		} else if item != " " && item != "" && item != "  " {
			e++
			recipt["e"] = e
		}
	}
	if recipt["e"] == 0 {
		delete(recipt, "e")
	}
	if recipt["\"\""] == 0 {
		delete(recipt, "\"\"")
	}
	return recipt
}
