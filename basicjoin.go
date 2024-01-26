package piscine

func BasicJoin(elems []string) string {
	var res string
	for _, words := range elems {
		res = res + words
	}
	return res
}
