package model

import (
	"github.com/huandu/go-clone"
	"time"
)

const CompanyHistoryStream = "company_history"

const (
	ObservableCompanyWbUserId           = WbUserId
	ObservableCompanyEmail              = Email
	ObservableCompanyAddress            = Address
	ObservableCompanyOrderLimit         = OrderLimit
	ObservableCompanyInn                = Inn
	ObservableCompanyCompanyName        = CompanyName
	ObservableCompanyFormOfOrganization = FormOfOrganization
	ObservableCompanyTaxSystem          = TaxForm
	ObservableCompanyCeoFio             = CeoFio
	ObservableCompanyIsActive           = IsActive
	ObservableCompanyIsTrusted          = IsTrusted
)

const (
	Inn                = "inn"
	FormOfOrganization = "form_of_organization"
	CompanyName        = "company_name"
	CeoFio             = "ceo_fio"
	WbUserId           = "wb_user_id"
	Address            = "address"
	OrderLimit         = "order_limit"
	IsActive           = "is_active"
	IsTrusted          = "is_trusted"
	Email              = "email"
	TaxForm            = "tax_form"
)

type Companies []Company

type Company struct {
	*Snap
	Id                 int64  `db:"id"`
	FormOfOrganization string `db:"form_of_organization"`
	Inn                string `db:"tax_id"`
	CompanyName        string `db:"company_name"`
	CeoFio             string `db:"ceo_fio"`
	CreatedDate        *time.Time
	UpdatedAt          *time.Time
	OrderLimit         int64  `db:"order_limit"`
	IsActive           bool   `db:"is_active"`
	IsTrusted          bool   `db:"is_trusted"`
	Address            string `db:"address"`
	Email              string `db:"email"`
	WbUserId           int64  `db:"wb_user_id"`
	TaxSystem          string `db:"tax_form"`
}

func (c *Company) GetTargetStream() string {
	return CompanyHistoryStream
}

func (c *Company) MakeNewSnapshot(tag string) {
	company := clone.Clone(c).(*Company)
	c.Snap.makeNewSnapshot(company, tag)
}
