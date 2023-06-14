package logging

type Event struct {
	Sequence 	uint64
	EventType 	EventType
	Key 		string
	Value 		string
}

type EventType byte
const (
	_ 					  = iota
	EventDelete EventType = iota
	EventPut              
)
