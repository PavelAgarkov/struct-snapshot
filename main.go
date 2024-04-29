package main

import (
	"fmt"
	"struct-snapshot/internal/event"
	"struct-snapshot/internal/event_dispatcher"
	"struct-snapshot/internal/model"
	"struct-snapshot/internal/subscriber"
	_ "time"
)

func main() {
	old := &model.Order{
		Snap: model.NewSnapshot(
			&model.Order{},
			&model.OrderHistory{},
		),
	}
	old.MakeNewSnapshot("v1")

	old.BrigadeName = "777"
	old.IsBlocked = true
	old.StatusInt = 33
	old.MakeNewSnapshot("v2")
	old.MakeNewSnapshot("v3")

	fmt.Println(old.GetSnapshots())
	//var lastkey int64
	for k, v := range old.GetSnapshots() {
		fmt.Println(k, v)
		for k1, v1 := range v.GetChanges() {
			//fmt.Println()
			fmt.Println(k1, v1)
		}
		fmt.Println("------------")
		//lastkey = int64(k)
	}

	fmt.Println(old.GetSnapshotByTag("v2").GetChanges()["status"], "<---------------------->")
	fmt.Println(old.GetSnapshotByIndex(1), "<nil----------------------nil>")

	// event dispatcher
	ed := event_dispatcher.NewEventDispatcher()
	ed.RegisterSubscriber(
		subscriber.NewOrderSubscriber(),
		[]event.ListeningEvent{
			event.OrderListeningEvent,
			//event.FunListeningEvent,
		},
	)

	ed.RegisterSubscriber(
		subscriber.NewFunSubscriber(),
		[]event.ListeningEvent{
			event.OrderListeningEvent,
			event.FunListeningEvent,
		},
	)

	ed.Dispatch(event.NewOrderEvent())
	ed.Dispatch(event.NewFunEvent())
}
