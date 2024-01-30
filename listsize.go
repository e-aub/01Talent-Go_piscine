package piscine

func ListSize(l *List) int {
	var count int
	temp := l.Head
	for temp != nil {
		count++
		temp = temp.Next
	}
	return count
}
