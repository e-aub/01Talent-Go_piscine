package piscine

func Split(s, sep string) []string {
	var result []string
	var sRes string
	for i := 0; i < len(s)-len(sep)+1; i++ {
		if s[i:len(sep)+i] == sep {
			if sRes != "" {
				result = append(result, sRes)
				sRes = ""
				i++
			}
		} else {
			sRes = sRes + string(s[i])
		}
		if i == len(s)-len(sep) {
			result = append(result, sRes)
		}
	}
	return result
}
