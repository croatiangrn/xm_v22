package repository

import (
	"context"
	"github.com/croatiangrn/xm_v22/internal/domain/company"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
)

var _ company.Repository = &CompanyRepository{}

type CompanyRepository struct {
	db *pgxpool.Pool
}

func NewCompanyRepository(db *pgxpool.Pool) company.Repository {
	return &CompanyRepository{db}
}

func (r *CompanyRepository) FindByID(ctx context.Context, id uuid.UUID) (*company.Company, error) {
	var c company.Company

	query := `SELECT id, name, description, amount_of_employees, registered, type FROM companies WHERE id = $1`

	if err := r.db.QueryRow(ctx, query, id).Scan(&c.ID, &c.Name); err != nil {
		return nil, err
	}

	return &c, nil
}

func (r *CompanyRepository) Create(ctx context.Context, c *company.Company) error {
	query := `INSERT INTO companies (id, name, description, amount_of_employees, registered, type) VALUES ($1, $2, $3, $4, $5, $6)`

	companyUUID, err := uuid.NewV7()
	if err != nil {
		return err
	}

	if _, err := r.db.Exec(ctx, query, companyUUID, c.Name, c.Description, c.AmountOfEmployees, c.Registered, c.Type); err != nil {
		return err
	}

	c.ID = companyUUID

	return nil
}

func (r *CompanyRepository) Update(ctx context.Context, obj *company.Company) error {
	query := `UPDATE companies SET name = $1, description = $2, amount_of_employees = $3, registered = $4, type = $5 WHERE id = $6`

	if _, err := r.db.Exec(ctx, query, obj.Name, obj.Description, obj.AmountOfEmployees, obj.Registered, obj.Type, obj.ID); err != nil {
		return err
	}

	return nil
}

func (r *CompanyRepository) Delete(ctx context.Context, id uuid.UUID) error {
	query := `DELETE FROM companies WHERE id = $1`

	if _, err := r.db.Exec(ctx, query, id); err != nil {
		return err
	}

	return nil
}
