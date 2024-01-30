package piscine

func ListAt(l *NodeL, pos int) *NodeL {
	Current := l
	for i := 0; i < pos && Current != nil; i++ {
		Current = Current.Next
	}
	return Current
}
