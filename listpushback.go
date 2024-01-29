package piscine

type NodeL1 struct {
	Data interface{}
	Next *NodeL1
}

type List1 struct {
	Head *NodeL1
	Tail *NodeL1
}

func ListPushBack(l *List1, data interface{}) {
	newNode := &NodeL1{Data: data, Next: nil}

	if l.Head == nil {
		l.Head = newNode
		l.Tail = newNode
	} else {
		l.Tail.Next = newNode
		l.Tail = newNode
	}
}
