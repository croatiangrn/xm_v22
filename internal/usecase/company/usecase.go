package company

import (
	"context"
	"github.com/croatiangrn/xm_v22/internal/controller/http/dto"
	"github.com/croatiangrn/xm_v22/internal/domain/company"
	"github.com/google/uuid"
)

type UseCase interface {
	GetCompany(ctx context.Context, id uuid.UUID) (*company.Company, error)
	CreateCompany(ctx context.Context, req dto.CreateCompanyRequest) (*company.Company, error)
	UpdateCompany(ctx context.Context, req dto.UpdateCompanyRequest, id uuid.UUID) (*company.Company, error)
	DeleteCompany(ctx context.Context, id uuid.UUID) error
}
