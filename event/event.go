package event

import "fmt"

type EventType int

const (
	ENone = EventType(iota)
	EButton
	EText
	ECheck
	ERadio
	ESelect
)

func Event(id string, e EventType, value string) {
	fmt.Println("new event:", id, e, value)
}
