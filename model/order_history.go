package model

type OrderHistory struct {
	OldEntity     *Order
	UpdatedEntity *Order
}

func (oh *OrderHistory) InitStructs() {
	oh.OldEntity = &Order{}
	oh.UpdatedEntity = &Order{}
}

func (oh *OrderHistory) ToFillEntitiesFields(oldEntity Entity, updatedEntity Entity) {
	if oldEntity != nil {
		order := oldEntity.(*Order)

		oh.OldEntity.CompanyName = order.CompanyName
		oh.OldEntity.BrigadeName = order.BrigadeName
		oh.OldEntity.StatusInt = order.StatusInt
		oh.OldEntity.IsBlocked = order.IsBlocked
	}
	if updatedEntity != nil {
		order := updatedEntity.(*Order)

		oh.UpdatedEntity.CompanyName = order.CompanyName
		oh.UpdatedEntity.BrigadeName = order.BrigadeName
		oh.UpdatedEntity.StatusInt = order.StatusInt
		oh.UpdatedEntity.IsBlocked = order.IsBlocked
	}
}

func (oh *OrderHistory) PrepareFieldsData() Changes {
	observableFields := make(map[string]*Field)
	if !oh.FieldIsEqual(oh.OldEntity.BrigadeName, oh.UpdatedEntity.BrigadeName) {
		observableFields[ObservableOrderBrigade] = &Field{
			OldValue: oh.OldEntity.BrigadeName,
			NewValue: oh.UpdatedEntity.BrigadeName,
		}
	}
	if !oh.FieldIsEqual(oh.OldEntity.CompanyName, oh.UpdatedEntity.CompanyName) {
		observableFields[ObservableOrderCompany] = &Field{
			OldValue: oh.OldEntity.CompanyName,
			NewValue: oh.UpdatedEntity.CompanyName,
		}
	}
	if !oh.FieldIsEqual(oh.OldEntity.StatusInt, oh.UpdatedEntity.StatusInt) {
		observableFields[ObservableOrderStatus] = &Field{
			OldValue: oh.OldEntity.StatusInt,
			NewValue: oh.UpdatedEntity.StatusInt,
		}
	}
	if !oh.FieldIsEqual(oh.OldEntity.IsBlocked, oh.UpdatedEntity.IsBlocked) {
		observableFields[ObservableOrderIsBlocked] = &Field{
			OldValue: oh.OldEntity.IsBlocked,
			NewValue: oh.UpdatedEntity.IsBlocked,
		}
	}

	return observableFields
}

func (oh *OrderHistory) FieldIsEqual(old any, new any) bool {
	return old == new
}

func (oh *OrderHistory) Reset() {
	oh.OldEntity = nil
	oh.UpdatedEntity = nil
}
