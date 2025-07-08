package reset_password_code_mongo_repository

import (
	"context"
	configs "ecommerce-white-label-backend/cmd/config"
	"ecommerce-white-label-backend/internal/domain/entity"
	domain_repository "ecommerce-white-label-backend/internal/domain/repository"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type ResetPasswordCodeRepository struct {
	db         *mongo.Database
	collection *mongo.Collection
}

func NewResetPasswordCodeRepository(db *mongo.Database) domain_repository.ResetPasswordCodeRepositoryInterface {
	collection := db.Collection(configs.MongoCfg.ResetPasswordCodeCollection)

	return &ResetPasswordCodeRepository{
		db:         db,
		collection: collection,
	}
}

func (f *ResetPasswordCodeRepository) Create(ctx context.Context, input entity.ResetPasswordCode) (int, error) {
	resetPasswordCodeEntity := ResetPasswordCodeModel{
		UserUuid:         input.UserUuid,
		Code:             input.Code,
		CodeExpiration:   time.Now().Add(15 * time.Minute).Format(time.RFC3339),
		AlreadyActivated: false,
		Email:            input.Email,
		CreatedAt:        time.Now().Format(time.RFC3339),
		ActivatedAt:      "",
	}

	_, err := f.collection.InsertOne(ctx, resetPasswordCodeEntity)

	if err != nil {
		return 0, err
	}

	return input.Code, nil
}

func (f *ResetPasswordCodeRepository) FindActive(ctx context.Context, email string) (int, error) {
	var resetPasswordCodeEntity ResetPasswordCodeModel
	err := f.collection.FindOne(ctx, bson.M{
		"email":             email,
		"already_activated": false,
		"code_expiration":   bson.M{"$gt": time.Now().Format(time.RFC3339)},
	}).Decode(&resetPasswordCodeEntity)

	if err != nil {
		return 0, err
	}

	return resetPasswordCodeEntity.Code, nil
}

func (f *ResetPasswordCodeRepository) IsValidCode(ctx context.Context, email string, code int) (bool, error) {
	var resetPasswordCodeEntity ResetPasswordCodeModel
	err := f.collection.FindOne(ctx, bson.M{
		"email": email,
		"code":  code,
	}).Decode(&resetPasswordCodeEntity)

	if err != nil {
		return false, err
	}

	return true, nil
}

func (f *ResetPasswordCodeRepository) ActivateCode(ctx context.Context, email string, code int) error {
	_, err := f.collection.UpdateOne(ctx, bson.M{
		"email": email,
		"code":  code,
	}, bson.D{
		{"$set", bson.M{"already_activated": true}},
		{"$set", bson.M{"activated_at": time.Now().Format(time.RFC3339)}},
	})

	return err
}
