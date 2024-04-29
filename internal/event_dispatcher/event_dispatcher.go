package event_dispatcher

import (
	"slices"
	"struct-snapshot/internal/event"
	"struct-snapshot/internal/subscriber"
)

type EventDispatcher struct {
	subscribers []subscriber.Subscriber
}

func NewEventDispatcher() *EventDispatcher {
	return &EventDispatcher{}
}

func (ed *EventDispatcher) RegisterSubscriber(subscriber subscriber.Subscriber, events []event.ListeningEvent) {
	subscriber.SetListenEvents(events)
	ed.subscribers = append(ed.subscribers, subscriber)
}

func (ed *EventDispatcher) Dispatch(event event.Event) {
	for _, sub := range ed.subscribers {
		if slices.Contains(sub.GetListenEvents(), event.GetName()) {
			sub.ListenTo(event)
		}
	}
}
