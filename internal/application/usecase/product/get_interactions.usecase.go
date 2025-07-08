package product_usecase

import (
	"context"
	configs "ecommerce-white-label-backend/cmd/config"
	storage_adapter "ecommerce-white-label-backend/internal/domain/adapters/storage"
	domain_repository "ecommerce-white-label-backend/internal/domain/repository"
	domain_response "ecommerce-white-label-backend/internal/domain/response"
	domain_product_usecase "ecommerce-white-label-backend/internal/domain/usecase/product"
	"time"
)

type GetProductInteractionUsecase struct {
	ProductRepository domain_repository.ProductRepositoryInterface
	UserRepository    domain_repository.UserRepositoryInterface
	StorageAdapter    storage_adapter.StorageAdapterInterface
}

func NewGetProductInteractionUsecase(
	repository domain_repository.ProductRepositoryInterface,
	userRepository domain_repository.UserRepositoryInterface,
	storageAdapter storage_adapter.StorageAdapterInterface,
) domain_product_usecase.GetInteractionUsecaseInterface {
	return &GetProductInteractionUsecase{
		ProductRepository: repository,
		UserRepository:    userRepository,
		StorageAdapter:    storageAdapter,
	}
}

func (u *GetProductInteractionUsecase) Execute(ctx context.Context, productUuid string) (map[string][]domain_response.ProductInteractionResponse, error) {
	groupedResponse := make(map[string][]domain_response.ProductInteractionResponse)

	repositoryResponse, err := u.ProductRepository.GetByUuid(ctx, productUuid)
	if err != nil {
		return nil, err
	}

	for _, interaction := range repositoryResponse.Reactions {
		user, err := u.UserRepository.GetByUuid(ctx, interaction.UserUuid)
		if err != nil {
			return nil, err
		}

		url, err := u.StorageAdapter.GeneratePresignedURL(
			ctx,
			configs.MinIoCfg.ProfileBucket,
			*user.Photo,
			time.Minute*15,
		)
		if err != nil {
			return nil, err
		}

		resp := domain_response.ProductInteractionResponse{
			UserUuid: user.Uuid,
			Photo:    url,
			Name:     user.Name,
			Emoji:    interaction.Emoji,
		}

		groupedResponse[interaction.Emoji] = append(groupedResponse[interaction.Emoji], resp)
	}

	return groupedResponse, nil
}
