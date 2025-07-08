package user_usecase

import (
	"context"
	configs "ecommerce-white-label-backend/cmd/config"
	storage_adapter "ecommerce-white-label-backend/internal/domain/adapters/storage"
	"ecommerce-white-label-backend/internal/domain/dto"
	domain_repository "ecommerce-white-label-backend/internal/domain/repository"
	domain_usecase "ecommerce-white-label-backend/internal/domain/usecase/user"
	"errors"
	"fmt"
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

	fmt.Println(input)

	if err == nil {
		if input.Photo != nil {
			err = u.StorageAdapter.UploadMedia(
				ctx,
				configs.MinIoCfg.ProfileBucket,
				"avatar/"+input.Uuid,
				input.Photo.File,
				input.Photo.FileSize,
				input.Photo.ContentType,
			)

			if err != nil {
				return errors.New("failed to upload user photo: " + err.Error())
			}
		}
	} else {
		return err
	}

	return err
}
