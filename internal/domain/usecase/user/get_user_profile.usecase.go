package domain_usecase

import (
	"context"
	"ecommerce-white-label-backend/internal/domain/dto"
	domain_response "ecommerce-white-label-backend/internal/domain/response"
)

type GetUserProfileUsecaseInterface interface {
	Execute(ctx context.Context, input dto.GetUserInputDto) (*domain_response.GetUserProfileResponse, error)
}
