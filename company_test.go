package struct_snapshot

import (
	"github.com/huandu/go-clone"
	"time"
)

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
	Id                 int64
	FormOfOrganization string
	Inn                string
	CompanyName        string
	CeoFio             string
	CreatedDate        *time.Time
	UpdatedAt          *time.Time
	OrderLimit         int64
	IsActive           bool
	IsTrusted          bool
	Address            string
	Email              string
	WbUserId           int64
	TaxSystem          string
}

func (c *Company) DoSnapshot(tag string) {
	companyClone := clone.Clone(c).(*Company)
	c.Snap.makeNewSnapshot(companyClone, tag)
}
