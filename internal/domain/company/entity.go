package company

type Company struct {
	ID                string `json:"id" gorm:"column:id;uuid;primaryKey;"`
	Name              string `json:"name" gorm:"column:id;varchar(15);unique;not null"`
	Description       string `json:"description" gorm:"column:description;varchar(3000)"`
	AmountOfEmployees int    `json:"amount_of_employees" gorm:"column:amount_of_employees;int;not null"`
	Registered        bool   `json:"registered" gorm:"column:registered;bool;not null"`
	Type              string `json:"type" gorm:"column:type;varchar(255);not null"`
}
