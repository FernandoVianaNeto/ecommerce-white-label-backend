package user_usecase

import (
	"context"
	"ecommerce-white-label-backend/internal/domain/dto"
	domain_repository "ecommerce-white-label-backend/internal/domain/repository"
	domain_response "ecommerce-white-label-backend/internal/domain/response"
	domain_usecase "ecommerce-white-label-backend/internal/domain/usecase/user"
)

type GetUserProfileUsecase struct {
	UserRepository domain_repository.UserRepositoryInterface
}

func NewGetUserProfileUseCase(
	repository domain_repository.UserRepositoryInterface,
) domain_usecase.GetUserProfileUsecaseInterface {
	return &GetUserProfileUsecase{
		UserRepository: repository,
	}
}

func (g *GetUserProfileUsecase) Execute(ctx context.Context, input dto.GetUserInputDto) (*domain_response.GetUserProfileResponse, error) {
	repositoryResponse, err := g.UserRepository.GetByUuid(ctx, input.Uuid)

	if repositoryResponse == nil {
		return nil, err
	}

	response := domain_response.GetUserProfileResponse{
		Uuid:            repositoryResponse.Uuid,
		Name:            repositoryResponse.Name,
		Email:           repositoryResponse.Email,
		ShippingAddress: repositoryResponse.ShippingAddress,
		BillingAddress:  repositoryResponse.BillingAddress,
		BirthDate:       repositoryResponse.BirthDate,
	}

	return &response, err
}
