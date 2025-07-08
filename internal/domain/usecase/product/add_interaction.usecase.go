package domain_product_usecase

import (
	"context"
	"ecommerce-white-label-backend/internal/domain/dto"
)

type AddInteractionUsecaseInterface interface {
	Execute(ctx context.Context, input dto.AddProductInteractionInputDto) error
}
