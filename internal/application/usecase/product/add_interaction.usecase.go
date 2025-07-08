package product_usecase

import (
	"context"
	"ecommerce-white-label-backend/internal/domain/dto"
	domain_repository "ecommerce-white-label-backend/internal/domain/repository"
	domain_product_usecase "ecommerce-white-label-backend/internal/domain/usecase/product"
	"errors"
	"time"
)

type AddInteractionUsecase struct {
	ProductRepository domain_repository.ProductRepositoryInterface
}

func NewAddInteractionUseCase(
	repository domain_repository.ProductRepositoryInterface,
) domain_product_usecase.AddInteractionUsecaseInterface {
	return &AddInteractionUsecase{
		ProductRepository: repository,
	}
}

func (u *AddInteractionUsecase) Execute(ctx context.Context, input dto.AddProductInteractionInputDto) error {
	value := ctx.Value("user_uuid")

	userUuid, ok := value.(string)

	if !ok {
		return errors.New("user uuid not found in context")
	}

	input.UserUuid = userUuid
	input.Timestamp = time.Now().Format(time.RFC3339)

	err := u.ProductRepository.AddInteraction(ctx, input)

	return err
}
