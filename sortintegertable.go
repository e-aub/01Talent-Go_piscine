package piscine

func SortIntegerTable(numbers []int) {
	for i := 0; i < len(numbers); i++ {
		for j := 0; j < len(numbers); j++ {
			if numbers[j] > numbers[i] {
				numbers[i], numbers[j] = numbers[j], numbers[i]
			}
		}
	}
}
