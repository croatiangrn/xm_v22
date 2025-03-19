package company

import (
	"context"
	"github.com/google/uuid"
)

type Repository interface {
	FindByID(ctx context.Context, id uuid.UUID) (*Company, error)
	Create(ctx context.Context, c *Company) error
	Update(ctx context.Context, obj *Company) error
	Delete(ctx context.Context, id uuid.UUID) error
}
