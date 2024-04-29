package event

var OrderListeningEvent ListeningEvent = "event.order"

func NewOrderEvent() *OrderEvent {
	return &OrderEvent{
		name: OrderListeningEvent,
	}
}

type OrderEvent struct {
	name ListeningEvent
}

func (oe *OrderEvent) GetName() ListeningEvent {
	return oe.name
}

func (oe *OrderEvent) GetData() string {
	return ""
}
