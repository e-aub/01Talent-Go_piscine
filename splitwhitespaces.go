package piscine

func SplitWhiteSpaces(s string) []string {
	var s2 string
	var sString []string
	for i := 0; i < len(s); i++ {
		if s[i] == ' ' || s[i] == '\n' || s[i] == '\t' {
			sString = append(sString, s2)
			s2 = ""
		} else {
			s2 = s2 + string(s[i])
		}
		if i == len(s)-1 {
			sString = append(sString, s2)
		}
	}
	return sString
}
