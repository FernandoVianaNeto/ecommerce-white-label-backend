package domain_product_usecase

import (
	"context"
)

type DeleteProductUsecaseInterface interface {
	Execute(ctx context.Context, workoutUuid string) error
}
