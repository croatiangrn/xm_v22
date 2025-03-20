package company

import (
	"context"
	"fmt"
	"github.com/croatiangrn/xm_v22/internal/controller/http/dto"
	"github.com/croatiangrn/xm_v22/internal/domain/company"
	"github.com/croatiangrn/xm_v22/internal/domain/event"
	"github.com/google/uuid"
)

var _ UseCase = &Interactor{}

type Interactor struct {
	repo     company.Repository
	producer event.ProducerInterface
}

func NewInteractor(repo company.Repository, producer event.ProducerInterface) *Interactor {
	return &Interactor{repo: repo, producer: producer}
}

func (uc *Interactor) GetCompany(ctx context.Context, id uuid.UUID) (*company.Company, error) {
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

	if err := uc.producer.Publish(ctx, "company-events", event.EventTypeCreateCompany, companyObj); err != nil {
		return nil, err
	}

	return companyObj, nil
}

func (uc *Interactor) UpdateCompany(ctx context.Context, req dto.UpdateCompanyRequest, id uuid.UUID) (*company.Company, error) {
	companyObj, err := uc.repo.FindByID(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("error finding company: %w", err)
	}

	if err := companyObj.AssignName(req.Name); err != nil {
		return nil, err
	}
	if err := companyObj.AssignDescription(req.Description); err != nil {
		return nil, err
	}

	if err := companyObj.AssignAmountOfEmployees(req.AmountOfEmployees); err != nil {
		return nil, err
	}

	companyObj.AssignRegistered(req.Registered)

	if err := companyObj.AssignType(req.Type); err != nil {
		return nil, err
	}

	if err := uc.repo.Update(ctx, companyObj); err != nil {
		return nil, err
	}

	if err := uc.producer.Publish(ctx, "company-events", event.EventTypeUpdateCompany, companyObj); err != nil {
		return nil, err
	}

	return companyObj, nil
}

func (uc *Interactor) DeleteCompany(ctx context.Context, id uuid.UUID) error {
	if err := uc.repo.Delete(ctx, id); err != nil {
		return fmt.Errorf("error deleting company: %w", err)
	}

	if err := uc.producer.Publish(ctx, "company-events", event.EventTypeDeleteCompany, &company.Company{ID: id}); err != nil {
		return err
	}

	return nil
}
