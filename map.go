package piscine

func Map(f func(int) bool, a []int) []bool {
	boolMap := make([]bool, len(a))
	for index, nb := range a {
		boolMap[index] = f(nb)
	}
	return boolMap
}
