package company

import "github.com/google/uuid"

type TypeOfCompany string

const (
	CompanyTypeCorporations       TypeOfCompany = "Corporations"
	CompanyTypeNonProfit          TypeOfCompany = "NonProfit"
	CompanyTypeCooperative        TypeOfCompany = "Cooperative"
	CompanyTypeSoleProprietorship TypeOfCompany = "Sole Proprietorship"
)

type Company struct {
	ID                uuid.UUID     `json:"id" gorm:"type:uuid;primaryKey"`
	Name              string        `json:"name" gorm:"column:name;varchar(15);uniqueIndex:uix_company_name;not null"`
	Description       string        `json:"description" gorm:"column:description;varchar(3000)"`
	AmountOfEmployees int           `json:"amount_of_employees" gorm:"column:amount_of_employees;int;not null"`
	Registered        bool          `json:"registered" gorm:"column:registered;bool;not null"`
	Type              TypeOfCompany `json:"type" gorm:"column:type;type:company_type;varchar(255);not null"`
}
