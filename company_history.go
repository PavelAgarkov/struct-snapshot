package main

type CompanyHistory struct {
	OldEntity     *Company
	UpdatedEntity *Company
}

func (ch *CompanyHistory) InitStructs() {
	ch.OldEntity = &Company{}
	ch.UpdatedEntity = &Company{}
}

func (ch *CompanyHistory) ToFillEntitiesFields(oldEntity Entity, updatedEntity Entity) {
	if oldEntity != nil {
		company := oldEntity.(*Company)

		ch.OldEntity.WbUserId = company.WbUserId
		ch.OldEntity.Email = company.Email
		ch.OldEntity.Address = company.Address
		ch.OldEntity.OrderLimit = company.OrderLimit
		ch.OldEntity.Inn = company.Inn
		ch.OldEntity.CompanyName = company.CompanyName
		ch.OldEntity.FormOfOrganization = company.FormOfOrganization
		ch.OldEntity.TaxSystem = company.TaxSystem
		ch.OldEntity.CeoFio = company.CeoFio
		ch.OldEntity.IsActive = company.IsActive
		ch.OldEntity.IsTrusted = company.IsTrusted
	}

	if updatedEntity != nil {
		company := updatedEntity.(*Company)

		ch.UpdatedEntity.WbUserId = company.WbUserId
		ch.UpdatedEntity.Email = company.Email
		ch.UpdatedEntity.Address = company.Address
		ch.UpdatedEntity.OrderLimit = company.OrderLimit
		ch.UpdatedEntity.Inn = company.Inn
		ch.UpdatedEntity.CompanyName = company.CompanyName
		ch.UpdatedEntity.FormOfOrganization = company.FormOfOrganization
		ch.UpdatedEntity.TaxSystem = company.TaxSystem
		ch.UpdatedEntity.CeoFio = company.CeoFio
		ch.UpdatedEntity.IsActive = company.IsActive
		ch.UpdatedEntity.IsTrusted = company.IsTrusted
	}
}

func (ch *CompanyHistory) PrepareFieldsData() Changes {
	observableFields := make(map[string]*Field)
	if !ch.FieldIsEqual(ch.OldEntity.WbUserId, ch.UpdatedEntity.WbUserId) {
		observableFields[ObservableCompanyWbUserId] = &Field{
			OldValue: ch.OldEntity.WbUserId,
			NewValue: ch.UpdatedEntity.WbUserId,
		}
	}
	if !ch.FieldIsEqual(ch.OldEntity.Email, ch.UpdatedEntity.Email) {
		observableFields[ObservableCompanyEmail] = &Field{
			OldValue: ch.OldEntity.Email,
			NewValue: ch.UpdatedEntity.Email,
		}
	}
	if !ch.FieldIsEqual(ch.OldEntity.Address, ch.UpdatedEntity.Address) {
		observableFields[ObservableCompanyAddress] = &Field{
			OldValue: ch.OldEntity.Address,
			NewValue: ch.UpdatedEntity.Address,
		}
	}
	if !ch.FieldIsEqual(ch.OldEntity.OrderLimit, ch.UpdatedEntity.OrderLimit) {
		observableFields[ObservableCompanyOrderLimit] = &Field{
			OldValue: ch.OldEntity.OrderLimit,
			NewValue: ch.UpdatedEntity.OrderLimit,
		}
	}
	if !ch.FieldIsEqual(ch.OldEntity.Inn, ch.UpdatedEntity.Inn) {
		observableFields[ObservableCompanyInn] = &Field{
			OldValue: ch.OldEntity.Inn,
			NewValue: ch.UpdatedEntity.Inn,
		}
	}
	if !ch.FieldIsEqual(ch.OldEntity.CompanyName, ch.UpdatedEntity.CompanyName) {
		observableFields[ObservableCompanyCompanyName] = &Field{
			OldValue: ch.OldEntity.CompanyName,
			NewValue: ch.UpdatedEntity.CompanyName,
		}
	}
	if !ch.FieldIsEqual(ch.OldEntity.FormOfOrganization, ch.UpdatedEntity.FormOfOrganization) {
		observableFields[ObservableCompanyFormOfOrganization] = &Field{
			OldValue: ch.OldEntity.FormOfOrganization,
			NewValue: ch.UpdatedEntity.FormOfOrganization,
		}
	}
	if !ch.FieldIsEqual(ch.OldEntity.TaxSystem, ch.UpdatedEntity.TaxSystem) {
		observableFields[ObservableCompanyTaxSystem] = &Field{
			OldValue: ch.OldEntity.TaxSystem,
			NewValue: ch.UpdatedEntity.TaxSystem,
		}
	}
	if !ch.FieldIsEqual(ch.OldEntity.CeoFio, ch.UpdatedEntity.CeoFio) {
		observableFields[ObservableCompanyCeoFio] = &Field{
			OldValue: ch.OldEntity.CeoFio,
			NewValue: ch.UpdatedEntity.CeoFio,
		}
	}
	if !ch.FieldIsEqual(ch.OldEntity.IsActive, ch.UpdatedEntity.IsActive) {
		observableFields[ObservableCompanyIsActive] = &Field{
			OldValue: ch.OldEntity.IsActive,
			NewValue: ch.UpdatedEntity.IsActive,
		}
	}
	if !ch.FieldIsEqual(ch.OldEntity.IsTrusted, ch.UpdatedEntity.IsTrusted) {
		observableFields[ObservableCompanyIsTrusted] = &Field{
			OldValue: ch.OldEntity.IsTrusted,
			NewValue: ch.UpdatedEntity.IsTrusted,
		}
	}

	return observableFields
}

func (ch *CompanyHistory) FieldIsEqual(old any, new any) bool {
	return old == new
}

func (ch *CompanyHistory) Reset() {
	ch.OldEntity = nil
	ch.UpdatedEntity = nil
}
