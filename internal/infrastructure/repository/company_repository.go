package repository

import (
	"context"
	"errors"
	"github.com/croatiangrn/xm_v22/internal/domain/company"
	customErrors "github.com/croatiangrn/xm_v22/internal/pkg/errors"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
	"time"
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

	query := `SELECT id, name, description, amount_of_employees, registered, type, created_at, updated_at FROM companies WHERE id = $1 AND deleted_at IS NULL`

	if err := r.db.QueryRow(ctx, query, id).Scan(&c.ID, &c.Name, &c.Description, &c.AmountOfEmployees, &c.Registered, &c.Type, &c.CreatedAt, &c.UpdatedAt); err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, customErrors.NewNotFoundError("Company", id.String())
		}

		return nil, err
	}

	return &c, nil
}

func (r *CompanyRepository) Create(ctx context.Context, c *company.Company) error {
	currentTime := time.Now().UTC()
	updatedAt := currentTime

	query := `INSERT INTO companies (id, name, description, amount_of_employees, registered, type, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6, $7, $8)`

	companyUUID, err := uuid.NewV7()
	if err != nil {
		return err
	}

	if _, err := r.db.Exec(ctx, query, companyUUID, c.Name, c.Description, c.AmountOfEmployees, c.Registered, c.Type, currentTime, updatedAt); err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) && pgErr.Code == "23505" {
			return customErrors.NewBadRequestError("name", "company with this name already exists")
		}

		return customErrors.NewInternalServerError("company create", err)
	}

	c.ID = companyUUID

	return nil
}

func (r *CompanyRepository) Update(ctx context.Context, obj *company.Company) error {
	query := `UPDATE companies SET name = $1, description = $2, amount_of_employees = $3, registered = $4, type = $5 WHERE id = $6`

	if _, err := r.db.Exec(ctx, query, obj.Name, obj.Description, obj.AmountOfEmployees, obj.Registered, obj.Type, obj.ID); err != nil {
		return customErrors.NewInternalServerError("company update", err)
	}

	return nil
}

func (r *CompanyRepository) Delete(ctx context.Context, id uuid.UUID) error {
	query := `DELETE FROM companies WHERE id = $1`

	if _, err := r.db.Exec(ctx, query, id); err != nil {
		return customErrors.NewInternalServerError("company delete", err)
	}

	return nil
}
