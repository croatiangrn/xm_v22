package company

import (
	"context"
	"github.com/croatiangrn/xm_v22/internal/controller/http/dto"
	"github.com/croatiangrn/xm_v22/internal/domain/company"
	"github.com/croatiangrn/xm_v22/internal/infrastructure/kafka"
)

var _ UseCase = &Interactor{}

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

func (uc *Interactor) CreateCompany(ctx context.Context, req dto.CreateCompanyRequest) (*company.Company, error) {
	companyObj := &company.Company{
		Name:              req.Name,
		Description:       req.Description,
		AmountOfEmployees: req.AmountOfEmployees,
		Registered:        req.Registered,
		Type:              req.Type,
	}

	if err := uc.repo.Create(ctx, companyObj); err != nil {
		return nil, err
	}

	if err := uc.producer.Publish(ctx, "company-events", kafka.EventTypeCreateCompany, companyObj); err != nil {
		return nil, err
	}

	return companyObj, nil
}

func (uc *Interactor) UpdateCompany(ctx context.Context, companyObj *company.Company, req dto.UpdateCompanyRequest) error {
	companyObj.Name = req.Name
	companyObj.Description = req.Description
	companyObj.AmountOfEmployees = req.AmountOfEmployees
	companyObj.Registered = req.Registered
	companyObj.Type = req.Type

	if err := uc.repo.Update(ctx, companyObj); err != nil {
		return err
	}

	if err := uc.producer.Publish(ctx, "company-events", kafka.EventTypeUpdateCompany, companyObj); err != nil {
		return err
	}

	return nil
}
