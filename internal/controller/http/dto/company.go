package dto

import "github.com/croatiangrn/xm_v22/internal/domain/company"

type CreateCompanyRequest struct {
	Name              string                `json:"name" binding:"required"`
	Description       string                `json:"description" binding:"omitempty"`
	AmountOfEmployees int                   `json:"amount_of_employees" binding:"required,gte=0"`
	Registered        bool                  `json:"registered" binding:"omitempty"`
	Type              company.TypeOfCompany `json:"type" binding:"required,oneof=NonProfit Corporations Cooperative 'Sole Proprietorship'"`
}

type UpdateCompanyRequest struct {
	Name              string                `json:"name" binding:"required"`
	Description       string                `json:"description" binding:"omitempty"`
	AmountOfEmployees int                   `json:"amount_of_employees" binding:"required,gte=0"`
	Registered        bool                  `json:"registered" binding:"omitempty"`
	Type              company.TypeOfCompany `json:"type" binding:"required,oneof=NonProfit Corporations Cooperative 'Sole Proprietorship'"`
}
