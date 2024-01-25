package main

import "github.com/01-edu/z01"

type Door struct {
	state string
}

func PrintStr(s string) {
	for _, r := range s {
		z01.PrintRune(r)
	}
}

func IsDoorClose(ptrDoor *Door) bool {
	PrintStr("is the Door closed ?\n")
	ptrDoor.state = "OPEN"
	return true
}

func IsDoorOpen(ptrDoor *Door) bool {
	PrintStr("is the Door closed ?\n")
	ptrDoor.state = "CLOSE"
	return true
}

func CloseDoor(ptrDoor *Door) string {
	PrintStr("Door Closing...\n")
	ptrDoor.state = "CLOSE"
	return ptrDoor.state
}

func OpenDoor(ptrDoor *Door) string {
	PrintStr("Door Opening...\n")
	ptrDoor.state = "OPEN"
	return ptrDoor.state
}

func main() {
	door := &Door{}

	OpenDoor(door)
	if IsDoorClose(door) {
		OpenDoor(door)
	}
	if IsDoorOpen(door) {
		CloseDoor(door)
	}
	if door.state == "OPEN" {
		CloseDoor(door)
	}
}
