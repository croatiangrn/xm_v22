package company

import (
	"context"
	"fmt"
	"github.com/croatiangrn/xm_v22/internal/controller/http/dto"
	"github.com/croatiangrn/xm_v22/internal/domain/company"
	"github.com/croatiangrn/xm_v22/internal/domain/event"
	customErrors "github.com/croatiangrn/xm_v22/internal/pkg/errors"
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
	companyObj, err := uc.repo.FindByID(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("error finding company: %w", err)
	}

	return companyObj, nil
}

// CreateCompany creates a new company
// We're returning *company.Company (DTO) because it's 1:1 mapping with the domain object
func (uc *Interactor) CreateCompany(ctx context.Context, req dto.CreateCompanyRequest) (*company.Company, error) {
	companyObj := &company.Company{}

	if err := companyObj.AssignName(req.Name); err != nil {
		return nil, customErrors.NewBadRequestError("name", err.Error())
	}

	if err := companyObj.AssignDescription(req.Description); err != nil {
		return nil, customErrors.NewBadRequestError("description", err.Error())
	}

	if err := companyObj.AssignAmountOfEmployees(req.AmountOfEmployees); err != nil {
		return nil, customErrors.NewBadRequestError("amount_of_employees", err.Error())
	}

	companyObj.AssignRegistered(req.Registered)

	if err := companyObj.AssignType(req.Type); err != nil {
		return nil, customErrors.NewBadRequestError("type", err.Error())
	}

	if err := uc.repo.Create(ctx, companyObj); err != nil {
		return nil, err
	}

	if err := uc.producer.Publish(ctx, "company-events", event.TypeCreateCompany, companyObj); err != nil {
		return nil, customErrors.NewInternalServerError("company creation", err)
	}

	return companyObj, nil
}

func (uc *Interactor) UpdateCompany(ctx context.Context, req dto.UpdateCompanyRequest, id uuid.UUID) (*company.Company, error) {
	companyObj, err := uc.repo.FindByID(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("error finding company: %w", err)
	}

	if err := companyObj.AssignName(req.Name); err != nil {
		return nil, customErrors.NewBadRequestError("name", err.Error())
	}
	if err := companyObj.AssignDescription(req.Description); err != nil {
		return nil, customErrors.NewBadRequestError("description", err.Error())
	}

	if err := companyObj.AssignAmountOfEmployees(req.AmountOfEmployees); err != nil {
		return nil, customErrors.NewBadRequestError("amount_of_employees", err.Error())
	}

	companyObj.AssignRegistered(req.Registered)

	if err := companyObj.AssignType(req.Type); err != nil {
		return nil, customErrors.NewBadRequestError("type", err.Error())
	}

	if err := uc.repo.Update(ctx, companyObj); err != nil {
		return nil, err
	}

	if err := uc.producer.Publish(ctx, "company-events", event.TypeUpdateCompany, companyObj); err != nil {
		return nil, customErrors.NewInternalServerError("company update", err)
	}

	return companyObj, nil
}

func (uc *Interactor) DeleteCompany(ctx context.Context, id uuid.UUID) error {
	if err := uc.repo.Delete(ctx, id); err != nil {
		return err
	}

	if err := uc.producer.Publish(ctx, "company-events", event.TypeDeleteCompany, &company.Company{ID: id}); err != nil {
		return customErrors.NewInternalServerError("company delete", err)
	}

	return nil
}
