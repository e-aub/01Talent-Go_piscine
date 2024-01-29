package piscine

func ListSize(l *List) int {
	var count int
	current := l.Head
	for current != nil {
		count++
		current = current.Next
	}
	return count
}
