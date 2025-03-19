package company

import (
	"context"
	"github.com/croatiangrn/xm_v22/internal/domain/company"
)

type Interactor struct {
	repo company.Repository
}

func NewInteractor(repo company.Repository) *Interactor {
	return &Interactor{repo: repo}
}

func (uc *Interactor) GetCompany(ctx context.Context, id string) (*company.Company, error) {
	return uc.repo.FindByID(ctx, id)
}
