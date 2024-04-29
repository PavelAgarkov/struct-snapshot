package model

import (
	"github.com/huandu/go-clone"
	"time"
)

const OrderHistoryStream = "order_history"

const (
	ObservableOrderBrigade   = BrigadeReq
	ObservableOrderCompany   = CompanyReq
	ObservableOrderStatus    = Status
	ObservableOrderIsBlocked = IsBlocked
)

const (
	Status     = "status"
	CompanyReq = "company"
	BrigadeReq = "brigade"
	IsBlocked  = "is_blocked"
)

type Order struct {
	*Snap
	Id            int64
	Status        string
	StatusInt     int64
	OrderType     int64
	StatusAlias   string
	CompanyId     int64
	BrigadeId     int64
	IsBlocked     bool
	CreatedDate   *time.Time
	UpdatedAt     *time.Time
	UpdatedWho    int64
	StartJobDate  *time.Time
	FinishJobDate *time.Time

	BrigadeName string
	CompanyName string
	Total       int64
}

func (o *Order) GetTargetStream() string {
	return OrderHistoryStream
}

func (o *Order) MakeNewSnapshot(tag string) {
	order := clone.Clone(o).(*Order)
	o.Snap.makeNewSnapshot(order, tag)
}
