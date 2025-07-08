package domain_auth_usecase

import (
	"context"
	"ecommerce-white-label-backend/internal/domain/dto"
	domain_response "ecommerce-white-label-backend/internal/domain/response"
)

type ValidateResetPasswordCodeUsecaseInterface interface {
	Execute(ctx context.Context, input dto.ValidateResetPasswordCodeInputDto) (domain_response.ValidateResetPasswordCodeResponse, error)
}
