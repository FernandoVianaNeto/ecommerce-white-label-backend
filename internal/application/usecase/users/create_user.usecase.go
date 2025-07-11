package user_usecase

import (
	"context"
	storage_adapter "ecommerce-white-label-backend/internal/domain/adapters/storage"
	"ecommerce-white-label-backend/internal/domain/dto"
	"ecommerce-white-label-backend/internal/domain/entity"
	domain_repository "ecommerce-white-label-backend/internal/domain/repository"
	domain_service "ecommerce-white-label-backend/internal/domain/service"
	domain_usecase "ecommerce-white-label-backend/internal/domain/usecase/user"
	"errors"

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

	userUuid := uuid.New().String()

	entity := entity.NewUser(
		userUuid,
		input.Email,
		input.BirthDate,
		input.Name,
		&encryptedPassword,
		"local",
		nil,
		input.ShippingAddress,
		input.BillingAddress,
	)

	err = u.UserRepository.Create(ctx, *entity)

	return err
}
