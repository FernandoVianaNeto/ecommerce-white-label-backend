package domain_product_usecase

import (
	"context"
	"ecommerce-white-label-backend/internal/domain/dto"
	"ecommerce-white-label-backend/internal/domain/entity"
)

type GetProductDetailsUsecaseInterface interface {
	Execute(ctx context.Context, input dto.GetProductDetailsInputDto) (*entity.Product, error)
}
