package subscriber

import (
	"fmt"
	"github.com/PavelAgarkov/struct-snapshot/internal/event"
)

type FunSubscriber struct {
	listenEvents []event.ListeningEvent
}

func NewFunSubscriber() *FunSubscriber {
	return &FunSubscriber{}
}

func (os *FunSubscriber) SetListenEvents(events []event.ListeningEvent) {
	os.listenEvents = events
}

func (os *FunSubscriber) GetListenEvents() []event.ListeningEvent {
	return os.listenEvents
}

func (os *FunSubscriber) ListenTo(event event.Event) {
	fmt.Println(event, "_fun")
	return
}
