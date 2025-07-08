package domain_auth_usecase

import (
	"context"
	"ecommerce-white-label-backend/internal/domain/dto"
)

type ResetPasswordUsecaseInterface interface {
	Execute(ctx context.Context, input dto.ResetPasswordInputDto) error
}
