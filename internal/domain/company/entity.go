package company

type companyType string

const (
	CompanyTypeCorporations       companyType = "corporations"
	CompanyTypeNonProfit          companyType = "non_profit"
	CompanyTypeCooperative        companyType = "cooperative"
	CompanyTypeSoleProprietorship companyType = "sole_proprietorship"
)

type Company struct {
	ID                string      `json:"id" gorm:"column:id;uuid;primaryKey;"`
	Name              string      `json:"name" gorm:"column:id;varchar(15);uniqueIndex:uix_company_name;not null"`
	Description       string      `json:"description" gorm:"column:description;varchar(3000)"`
	AmountOfEmployees int         `json:"amount_of_employees" gorm:"column:amount_of_employees;int;not null"`
	Registered        bool        `json:"registered" gorm:"column:registered;bool;not null"`
	Type              companyType `json:"type" gorm:"column:type;type:company_type;varchar(255);not null"`
}
