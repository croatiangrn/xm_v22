package company

import "context"

type Repository interface {
	FindByID(ctx context.Context, id string) (*Company, error)
}
