package piscine

func Sortint(numbers []int) []int {
	for i := 0; i < len(numbers); i++ {
		for j := 0; j < len(numbers)-1; j++ {
			if numbers[j] > numbers[i] {
				numbers[i], numbers[j] = numbers[j], numbers[i]
			}
		}
	}
	return numbers
}
