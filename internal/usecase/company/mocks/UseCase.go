// Code generated by mockery v2.53.3. DO NOT EDIT.

package mocks

import (
	context "context"

	dto "github.com/croatiangrn/xm_v22/internal/controller/http/dto"
	mock "github.com/stretchr/testify/mock"

	uuid "github.com/google/uuid"
)

// UseCase is an autogenerated mock type for the UseCase type
type UseCase struct {
	mock.Mock
}

// CreateCompany provides a mock function with given fields: ctx, req
func (_m *UseCase) CreateCompany(ctx context.Context, req dto.CreateCompanyRequest) (*dto.CompanyResponse, error) {
	ret := _m.Called(ctx, req)

	if len(ret) == 0 {
		panic("no return value specified for CreateCompany")
	}

	var r0 *dto.CompanyResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, dto.CreateCompanyRequest) (*dto.CompanyResponse, error)); ok {
		return rf(ctx, req)
	}
	if rf, ok := ret.Get(0).(func(context.Context, dto.CreateCompanyRequest) *dto.CompanyResponse); ok {
		r0 = rf(ctx, req)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*dto.CompanyResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, dto.CreateCompanyRequest) error); ok {
		r1 = rf(ctx, req)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteCompany provides a mock function with given fields: ctx, id
func (_m *UseCase) DeleteCompany(ctx context.Context, id uuid.UUID) error {
	ret := _m.Called(ctx, id)

	if len(ret) == 0 {
		panic("no return value specified for DeleteCompany")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, uuid.UUID) error); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetCompany provides a mock function with given fields: ctx, id
func (_m *UseCase) GetCompany(ctx context.Context, id uuid.UUID) (*dto.CompanyResponse, error) {
	ret := _m.Called(ctx, id)

	if len(ret) == 0 {
		panic("no return value specified for GetCompany")
	}

	var r0 *dto.CompanyResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, uuid.UUID) (*dto.CompanyResponse, error)); ok {
		return rf(ctx, id)
	}
	if rf, ok := ret.Get(0).(func(context.Context, uuid.UUID) *dto.CompanyResponse); ok {
		r0 = rf(ctx, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*dto.CompanyResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, uuid.UUID) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateCompany provides a mock function with given fields: ctx, req, id
func (_m *UseCase) UpdateCompany(ctx context.Context, req dto.UpdatePatchCompanyRequest, id uuid.UUID) (*dto.CompanyResponse, error) {
	ret := _m.Called(ctx, req, id)

	if len(ret) == 0 {
		panic("no return value specified for UpdateCompany")
	}

	var r0 *dto.CompanyResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, dto.UpdatePatchCompanyRequest, uuid.UUID) (*dto.CompanyResponse, error)); ok {
		return rf(ctx, req, id)
	}
	if rf, ok := ret.Get(0).(func(context.Context, dto.UpdatePatchCompanyRequest, uuid.UUID) *dto.CompanyResponse); ok {
		r0 = rf(ctx, req, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*dto.CompanyResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, dto.UpdatePatchCompanyRequest, uuid.UUID) error); ok {
		r1 = rf(ctx, req, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewUseCase creates a new instance of UseCase. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewUseCase(t interface {
	mock.TestingT
	Cleanup(func())
}) *UseCase {
	mock := &UseCase{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
