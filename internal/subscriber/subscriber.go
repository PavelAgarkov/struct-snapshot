package subscriber

import "struct-snapshot/internal/event"

type Subscriber interface {
	SetListenEvents(events []event.ListeningEvent)
	GetListenEvents() []event.ListeningEvent
	ListenTo(event event.Event)
}
