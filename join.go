package piscine

func Join(strs []string, sep string) string {
	var s string
	for index, word := range strs {
		if index != len(strs)-1 {
			s = s + word + sep
		} else {
			s = s + word
		}
	}
	return s
}
