package dto

type CreateCompanyRequest struct {
	Name              string `json:"name"`
	Description       string `json:"description"`
	AmountOfEmployees int    `json:"amount_of_employees"`
	Registered        bool   `json:"registered"`
	Type              string `json:"type"`
}

type UpdateCompanyRequest struct {
	Name              string `json:"name"`
	Description       string `json:"description"`
	AmountOfEmployees int    `json:"amount_of_employees"`
	Registered        bool   `json:"registered"`
	Type              string `json:"type"`
}

type CompanyResponse struct {
	ID                string `json:"id"`
	Name              string `json:"name"`
	Description       string `json:"description"`
	AmountOfEmployees int    `json:"amount_of_employees"`
	Registered        bool   `json:"registered"`
	Type              string `json:"type"`
}
