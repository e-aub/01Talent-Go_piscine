package piscine

type NodeL struct {
	Data interface{}
	Next *NodeL
}

type List struct {
	Head *NodeL
	Tail *NodeL
}

func ListSize(l *List) int {
	var count int
	temp := l.Head
	for temp != nil {
		count++
		temp = temp.Next
	}
	return count
}
