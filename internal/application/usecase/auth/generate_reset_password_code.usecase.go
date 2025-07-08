package auth_usecase

import (
	"context"
	adapter "ecommerce-white-label-backend/internal/domain/adapters/email_sender"
	"ecommerce-white-label-backend/internal/domain/dto"
	"ecommerce-white-label-backend/internal/domain/entity"
	domain_repository "ecommerce-white-label-backend/internal/domain/repository"
	domain_usecase "ecommerce-white-label-backend/internal/domain/usecase/auth"
	"errors"
)

type GenerateResetPasswordCodeUsecase struct {
	ResetPasswordCodeRepository domain_repository.ResetPasswordCodeRepositoryInterface
	UserRepository              domain_repository.UserRepositoryInterface
	EmailSenderAdapter          adapter.EmailSenderAdapterInterface
}

func NewGenerateResetPasswordCodeUsecase(
	resetPasswordCodeRepository domain_repository.ResetPasswordCodeRepositoryInterface,
	userRepository domain_repository.UserRepositoryInterface,
	emailSenderAdapter adapter.EmailSenderAdapterInterface,
) domain_usecase.GenerateResetPasswordCodeUsecaseInterface {
	return &GenerateResetPasswordCodeUsecase{
		ResetPasswordCodeRepository: resetPasswordCodeRepository,
		UserRepository:              userRepository,
		EmailSenderAdapter:          emailSenderAdapter,
	}
}

func (a *GenerateResetPasswordCodeUsecase) Execute(ctx context.Context, input dto.GenerateResetPasswordCodeInputDto) error {
	user, err := a.UserRepository.GetByEmailAndAuthProvider(ctx, input.Email, "local")

	if err != nil {
		return err
	}

	if user == nil {
		return errors.New("user not found")
	}

	activeCode, _ := a.ResetPasswordCodeRepository.FindActive(ctx, input.Email)

	if activeCode != 0 {
		return nil
	}

	resetPasswordEntity, err := entity.NewResetPasswordCode(user.Uuid, user.Email)

	if err != nil {
		return err
	}

	newCode, err := a.ResetPasswordCodeRepository.Create(ctx, *resetPasswordEntity)

	if err != nil {
		return err
	}

	err = a.EmailSenderAdapter.SendResetPasswordEmail(ctx, user.Email, newCode)

	return err
}
