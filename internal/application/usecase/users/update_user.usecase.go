package user_usecase

import (
	"context"
	storage_adapter "ecommerce-white-label-backend/internal/domain/adapters/storage"
	"ecommerce-white-label-backend/internal/domain/dto"
	domain_repository "ecommerce-white-label-backend/internal/domain/repository"
	domain_usecase "ecommerce-white-label-backend/internal/domain/usecase/user"
)

type UpdateUserUsecase struct {
	UserRepository domain_repository.UserRepositoryInterface
	StorageAdapter storage_adapter.StorageAdapterInterface
}

func NewUpdateUserUseCase(
	repository domain_repository.UserRepositoryInterface,
	storage storage_adapter.StorageAdapterInterface,
) domain_usecase.UpdateUserUsecaseInterface {
	return &UpdateUserUsecase{
		UserRepository: repository,
		StorageAdapter: storage,
	}
}

func (u *UpdateUserUsecase) Execute(ctx context.Context, input dto.UpdateUserInputDto) error {
	err := u.UserRepository.UpdateByUuid(ctx, input)

	return err
}
