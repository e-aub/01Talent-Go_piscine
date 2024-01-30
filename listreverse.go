package piscine

func ListReverse(l *List) {
	var prv *NodeL
	curr := l.Head
	var nxt *NodeL
	for curr != nil {
		nxt = curr.Next
		curr.Next = prv
		prv = curr
		curr = nxt
	}
	l.Tail = l.Head
	l.Head = prv
}
