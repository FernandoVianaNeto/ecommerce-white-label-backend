package domain_usecase

import (
	"context"
	"ecommerce-white-label-backend/internal/domain/dto"
)

type UpdateUserUsecaseInterface interface {
	Execute(ctx context.Context, input dto.UpdateUserInputDto) error
}
