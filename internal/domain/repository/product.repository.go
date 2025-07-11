package domain_repository

import (
	"context"
	"ecommerce-white-label-backend/internal/domain/entity"
	domain_response "ecommerce-white-label-backend/internal/domain/response"
)

//go:generate mockgen -source $GOFILE -package $GOPACKAGE -destination $ROOT_DIR/test/mocks/$GOPACKAGE/mock_$GOFILE

const ProductCollection = "products"

type ProductRepositoryInterface interface {
	Create(ctx context.Context, entity entity.Product) error
	ListProducts(ctx context.Context, page string) (domain_response.ListProductsPaginatedResponse, error)
	GetByUuid(ctx context.Context, uuid string) (*entity.Product, error)
	// AddInteraction(ctx context.Context, input dto.AddProductInteractionInputDto) error
	// Delete(ctx context.Context, productUuid string) error
}
