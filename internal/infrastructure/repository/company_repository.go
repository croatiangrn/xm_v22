package repository

import (
	"context"
	"github.com/croatiangrn/xm_v22/internal/domain/company"
	"github.com/jackc/pgx/v5/pgxpool"
)

type CompanyRepository struct {
	db *pgxpool.Pool
}

func NewCompanyRepository(db *pgxpool.Pool) company.Repository {
	return &CompanyRepository{db}
}

func (r *CompanyRepository) FindByID(ctx context.Context, id string) (*company.Company, error) {
	var c company.Company

	query := `SELECT id, name, description, amount_of_employees, registered, type FROM companies WHERE id = $1`

	if err := r.db.QueryRow(ctx, query, id).Scan(&c.ID, &c.Name); err != nil {
		return nil, err
	}

	return &c, nil
}
