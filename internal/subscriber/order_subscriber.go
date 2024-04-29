package subscriber

import (
	"fmt"
	"struct-snapshot/internal/event"
)

type OrderSubscriber struct {
	listenEvents []event.ListeningEvent
	// тут моно запиисать указатель на сервис и вызвать его в Subscribe()
}

func NewOrderSubscriber() *OrderSubscriber {
	return &OrderSubscriber{}
}

func (os *OrderSubscriber) SetListenEvents(events []event.ListeningEvent) {
	os.listenEvents = events
}

func (os *OrderSubscriber) GetListenEvents() []event.ListeningEvent {
	return os.listenEvents
}

func (os *OrderSubscriber) ListenTo(event event.Event) {
	fmt.Println(event, "_order")
	return
}
