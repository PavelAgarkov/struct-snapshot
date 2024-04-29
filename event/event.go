package event

type ListeningEvent string

type Event interface {
	GetName() ListeningEvent
	GetData() string
}
