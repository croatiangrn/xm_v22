package models

import "time"

type Company struct {
	ID                string     `json:"id"`
	CreatedAt         time.Time  `json:"created_at"`
	UpdatedAt         time.Time  `json:"updated_at"`
	DeletedAt         *time.Time `json:"deleted_at"`
	Name              string     `json:"name"`
	Description       string     `json:"description"`
	AmountOfEmployees int        `json:"amount_of_employees"`
	Registered        bool       `json:"registered"`
	Type              string     `json:"type"`
}
