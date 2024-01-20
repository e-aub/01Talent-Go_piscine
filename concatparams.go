package piscine

func ConcatParams(args []string) string {
	var res string
	for index, arg := range args {
		if index == 0 {
			res = res + arg
		} else {
			res = res + "\n" + arg
		}
	}
	return res
}
