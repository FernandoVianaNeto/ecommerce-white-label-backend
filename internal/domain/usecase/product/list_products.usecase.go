package domain_product_usecase

import (
	"context"
	domain_response "ecommerce-white-label-backend/internal/domain/response"
)

type ListProductsUsecaseInterface interface {
	Execute(ctx context.Context, page string) (domain_response.ListProductsPaginatedResponse, error)
}
