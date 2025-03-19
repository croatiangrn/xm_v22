package company

import (
	"context"
	"github.com/croatiangrn/xm_v22/internal/domain/company"
)

type UseCase interface {
	GetCompany(ctx context.Context, id string) (*company.Company, error)
}
