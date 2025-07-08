package domain_auth_usecase

import (
	"context"
	"ecommerce-white-label-backend/internal/domain/dto"
	domain_response "ecommerce-white-label-backend/internal/domain/response"
)

type GoogleAuthUsecaseInterface interface {
	Execute(ctx context.Context, input dto.GoogleAuthInputDto) (domain_response.AuthResponse, error)
}
