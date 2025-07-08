package domain_usecase

import (
	"context"
	"ecommerce-white-label-backend/internal/domain/dto"
)

type CreateUserUsecaseInterface interface {
	Execute(ctx context.Context, input dto.CreateUserInputDto) error
}
