package company

import (
	"context"
	"github.com/croatiangrn/xm_v22/internal/domain/company"
	"github.com/croatiangrn/xm_v22/internal/infrastructure/kafka"
)

type Interactor struct {
	repo     company.Repository
	producer *kafka.Producer
}

func NewInteractor(repo company.Repository, producer *kafka.Producer) *Interactor {
	return &Interactor{repo: repo, producer: producer}
}

func (uc *Interactor) GetCompany(ctx context.Context, id string) (*company.Company, error) {
	return uc.repo.FindByID(ctx, id)
}
