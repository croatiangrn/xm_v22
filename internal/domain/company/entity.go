package company

import (
	"errors"
	"fmt"
	customErrors "github.com/croatiangrn/xm_v22/internal/pkg/errors"
	"github.com/google/uuid"
	"time"
)

const (
	maxDescriptionLength = 3000
)

var (
	companyTypes = []string{"NonProfit", "Corporations", "Cooperative", "Sole Proprietorship"}
)

type Company struct {
	ID                uuid.UUID
	Name              string
	Description       string
	AmountOfEmployees int
	Registered        bool
	Type              string
	CreatedAt         time.Time
	UpdatedAt         time.Time
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
		return customErrors.NewBadRequestError("amount_of_employees", "amount of employees cannot be negative")
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
