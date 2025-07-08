package domain_auth_usecase

import (
	"context"
	"ecommerce-white-label-backend/internal/domain/dto"
)

type GenerateResetPasswordCodeUsecaseInterface interface {
	Execute(ctx context.Context, input dto.GenerateResetPasswordCodeInputDto) error
}
