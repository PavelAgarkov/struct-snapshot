package subscriber

import (
	"github.com/PavelAgarkov/struct-snapshot/event"
)

type Subscriber interface {
	SetListenEvents(events []event.ListeningEvent)
	GetListenEvents() []event.ListeningEvent
	ListenTo(event event.Event)
}
