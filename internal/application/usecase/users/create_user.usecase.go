package user_usecase

import (
	"context"
	configs "ecommerce-white-label-backend/cmd/config"
	storage_adapter "ecommerce-white-label-backend/internal/domain/adapters/storage"
	"ecommerce-white-label-backend/internal/domain/dto"
	"ecommerce-white-label-backend/internal/domain/entity"
	domain_repository "ecommerce-white-label-backend/internal/domain/repository"
	domain_service "ecommerce-white-label-backend/internal/domain/service"
	domain_usecase "ecommerce-white-label-backend/internal/domain/usecase/user"
	"errors"
	"fmt"

	"github.com/google/uuid"
)

type CreateUserUsecase struct {
	UserRepository       domain_repository.UserRepositoryInterface
	EncryptStringService domain_service.EncryptStringServiceInterface
	StorageAdapter       storage_adapter.StorageAdapterInterface
}

func NewCreateUserUseCase(
	repository domain_repository.UserRepositoryInterface,
	encryptStringService domain_service.EncryptStringServiceInterface,
	storageAdapter storage_adapter.StorageAdapterInterface,
) domain_usecase.CreateUserUsecaseInterface {
	return &CreateUserUsecase{
		UserRepository:       repository,
		EncryptStringService: encryptStringService,
		StorageAdapter:       storageAdapter,
	}
}

func (u *CreateUserUsecase) Execute(ctx context.Context, input dto.CreateUserInputDto) error {
	var (
		encryptedPassword []byte
		err               error
	)

	user, err := u.UserRepository.GetByEmailAndAuthProvider(ctx, input.Email, input.Origin)

	if err != nil {
		return err
	}

	if user != nil {
		return errors.New("user already exists")
	}

	if input.Origin == "local" && input.Password != nil {
		encryptedPassword, err = u.EncryptStringService.EncryptString(*input.Password, 10)
		if err != nil {
			return err
		}
	}

	photoPath := "avatar/default"

	userUuid := uuid.New().String()

	if input.Photo != nil {
		photoPath = fmt.Sprintf("avatar/%s", userUuid)
	}

	entity := entity.NewUser(
		userUuid,
		input.Email,
		input.BirthDate,
		input.Name,
		&encryptedPassword,
		input.Sports,
		input.Origin,
		nil,
		&photoPath,
	)

	if input.Photo != nil {
		err = u.StorageAdapter.UploadMedia(
			ctx,
			configs.MinIoCfg.ProfileBucket,
			photoPath,
			input.Photo.File,
			input.Photo.FileSize,
			input.Photo.ContentType,
		)
		if err != nil {
			return err
		}
	}

	err = u.UserRepository.Create(ctx, *entity)

	return err
}
