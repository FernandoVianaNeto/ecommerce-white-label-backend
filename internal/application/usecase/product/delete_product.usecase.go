package product_usecase

import (
	"context"
	domain_repository "ecommerce-white-label-backend/internal/domain/repository"
	domain_product_usecase "ecommerce-white-label-backend/internal/domain/usecase/product"
)

type DeleteProductUsecase struct {
	ProductRepository domain_repository.ProductRepositoryInterface
}

func NewDeleteProductUseCase(
	repository domain_repository.ProductRepositoryInterface,
) domain_product_usecase.DeleteProductUsecaseInterface {
	return &DeleteProductUsecase{
		ProductRepository: repository,
	}
}

func (u *DeleteProductUsecase) Execute(ctx context.Context, productUuid string) error {
	// fmt.Println(productUuid)
	// err := u.ProductRepository.Delete(ctx, productUuid)

	return nil
}
