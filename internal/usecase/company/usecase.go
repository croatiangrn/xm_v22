package company

import (
	"context"
	"github.com/croatiangrn/xm_v22/internal/controller/http/dto"
	"github.com/croatiangrn/xm_v22/internal/domain/company"
)

type UseCase interface {
	GetCompany(ctx context.Context, id string) (*company.Company, error)
	CreateCompany(ctx context.Context, req dto.CreateCompanyRequest) (*company.Company, error)
	UpdateCompany(ctx context.Context, obj *company.Company, req dto.UpdateCompanyRequest) error
}
