package piscine

func StrRev(s string) string {
	var res string
	for _, char := range s {
		res = string(char) + res
	}
	return res
}
