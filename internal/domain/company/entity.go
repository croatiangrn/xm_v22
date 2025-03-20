package company

import (
	"errors"
	"fmt"
	"github.com/google/uuid"
)

const (
	maxDescriptionLength = 3000
)

var (
	companyTypes = []string{"NonProfit", "Corporations", "Cooperative", "Sole Proprietorship"}
)

type Company struct {
	ID                uuid.UUID `json:"id"`
	Name              string    `json:"name"`
	Description       string    `json:"description"`
	AmountOfEmployees int       `json:"amount_of_employees"`
	Registered        bool      `json:"registered"`
	Type              string    `json:"type"`
}

func (c *Company) AssignName(name string) error {
	if name == "" {
		return errors.New("name cannot be empty")
	}

	c.Name = name
	return nil
}

func (c *Company) AssignDescription(description string) error {
	if len(description) > maxDescriptionLength {
		return fmt.Errorf("description is too long, max length is %d", maxDescriptionLength)
	}

	c.Description = description
	return nil
}

func (c *Company) AssignAmountOfEmployees(amount int) error {
	if amount < 0 {
		return errors.New("amount of employees cannot be negative")
	}

	c.AmountOfEmployees = amount
	return nil
}

func (c *Company) AssignRegistered(registered bool) {
	c.Registered = registered
}

func (c *Company) AssignType(t string) error {
	if t == "" {
		return errors.New("type cannot be empty")
	}

	var found bool
	for _, v := range companyTypes {
		if v == t {
			found = true
			break
		}
	}

	if !found {
		return fmt.Errorf("type must be one of %v", companyTypes)
	}

	c.Type = t
	return nil
}
