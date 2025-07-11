package product_usecase

import (
	"context"
	domain_repository "ecommerce-white-label-backend/internal/domain/repository"
	domain_response "ecommerce-white-label-backend/internal/domain/response"
	domain_product_usecase "ecommerce-white-label-backend/internal/domain/usecase/product"
	"fmt"
)

type ListProductsUsecase struct {
	ProductRepository domain_repository.ProductRepositoryInterface
}

func NewListProductsUseCase(
	repository domain_repository.ProductRepositoryInterface,
) domain_product_usecase.ListProductsUsecaseInterface {
	return &ListProductsUsecase{
		ProductRepository: repository,
	}
}

func (u *ListProductsUsecase) Execute(ctx context.Context, page string) (domain_response.ListProductsPaginatedResponse, error) {
	// value := ctx.Value("user_uuid")

	// userUuid, ok := value.(string)

	// if !ok {
	// 	return nil, errors.New("user uuid not found in context")
	// }
	//TODO VALIDATE USE LOGIN
	fmt.Println("CHEGUEI AQUI")
	response, err := u.ProductRepository.ListProducts(ctx, page)

	return response, err
}
