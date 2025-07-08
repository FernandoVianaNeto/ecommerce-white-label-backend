package domain_product_usecase

import (
	"context"
	"ecommerce-white-label-backend/internal/domain/dto"
)

type CreateProductUsecaseInterface interface {
	Execute(ctx context.Context, input dto.CreateProductInputDto) error
}
