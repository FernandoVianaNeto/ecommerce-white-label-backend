package domain_repository

//go:generate mockgen -source $GOFILE -package $GOPACKAGE -destination $ROOT_DIR/test/mocks/$GOPACKAGE/mock_$GOFILE

import (
	"context"
	"ecommerce-white-label-backend/internal/domain/dto"
	"ecommerce-white-label-backend/internal/domain/entity"
)

const UserCollection = "users"

type UserRepositoryInterface interface {
	Create(ctx context.Context, entity entity.User) error
	GetByUuid(ctx context.Context, userUuid string) (*entity.User, error)
	GetByEmailAndAuthProvider(ctx context.Context, email string, authProvider string) (*entity.User, error)
	UpdateByUuid(ctx context.Context, input dto.UpdateUserInputDto) error
	UpdatePassword(ctx context.Context, input dto.UserResetPasswordInputDto) error
}
