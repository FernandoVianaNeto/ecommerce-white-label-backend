package user_usecase

import (
	"context"
	configs "ecommerce-white-label-backend/cmd/config"
	storage_adapter "ecommerce-white-label-backend/internal/domain/adapters/storage"
	"ecommerce-white-label-backend/internal/domain/dto"
	domain_repository "ecommerce-white-label-backend/internal/domain/repository"
	domain_response "ecommerce-white-label-backend/internal/domain/response"
	domain_usecase "ecommerce-white-label-backend/internal/domain/usecase/user"
	"time"
)

type GetUserProfileUsecase struct {
	UserRepository domain_repository.UserRepositoryInterface
	StorageAdapter storage_adapter.StorageAdapterInterface
}

func NewGetUserProfileUseCase(
	repository domain_repository.UserRepositoryInterface,
	storageAdapter storage_adapter.StorageAdapterInterface,
) domain_usecase.GetUserProfileUsecaseInterface {
	return &GetUserProfileUsecase{
		UserRepository: repository,
		StorageAdapter: storageAdapter,
	}
}

func (g *GetUserProfileUsecase) Execute(ctx context.Context, input dto.GetUserInputDto) (*domain_response.GetUserProfileResponse, error) {
	repositoryResponse, err := g.UserRepository.GetByUuid(ctx, input.Uuid)

	if repositoryResponse == nil {
		return nil, err
	}

	url, err := g.StorageAdapter.GeneratePresignedURL(
		ctx,
		configs.MinIoCfg.ProfileBucket,
		*repositoryResponse.Photo,
		time.Minute*15,
	)

	response := domain_response.GetUserProfileResponse{
		Name:      repositoryResponse.Name,
		Sports:    &repositoryResponse.Sports,
		Email:     repositoryResponse.Email,
		BirthDate: repositoryResponse.BirthDate,
		PhotoUrl:  url,
	}

	return &response, err
}
