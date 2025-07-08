package product_usecase

import (
	"context"
	configs "ecommerce-white-label-backend/cmd/config"
	"ecommerce-white-label-backend/internal/domain/adapters/messaging"
	storage_adapter "ecommerce-white-label-backend/internal/domain/adapters/storage"
	"ecommerce-white-label-backend/internal/domain/dto"
	"ecommerce-white-label-backend/internal/domain/entity"
	domain_repository "ecommerce-white-label-backend/internal/domain/repository"
	domain_product_usecase "ecommerce-white-label-backend/internal/domain/usecase/product"
	"encoding/json"
	"errors"
	"fmt"
	"log"

	"github.com/google/uuid"
)

type CreateProductUsecase struct {
	ProductRepository domain_repository.ProductRepositoryInterface
	StorageAdapter    storage_adapter.StorageAdapterInterface
	Messaging         messaging.Client
}

func NewCreateProductUseCase(
	repository domain_repository.ProductRepositoryInterface,
	storageAdapter storage_adapter.StorageAdapterInterface,
	messagingClient messaging.Client,
) domain_product_usecase.CreateProductUsecaseInterface {
	return &CreateProductUsecase{
		ProductRepository: repository,
		StorageAdapter:    storageAdapter,
		Messaging:         messagingClient,
	}
}

func (u *CreateProductUsecase) Execute(ctx context.Context, input dto.CreateProductInputDto) error {
	var (
		pace string
		err  error
	)

	value := ctx.Value("user_uuid")

	userUuid, ok := value.(string)

	if !ok {
		return errors.New("user uuid not found in context")
	}

	photoPath := "photo/default"

	ProductUuid := uuid.New().String()

	if input.Photo != nil {
		photoPath = fmt.Sprintf("photo/%s", userUuid)
	}

	entity := entity.NewProduct(
		ProductUuid,
		string(userUuid),
		input.Title,
		entity.Location(input.Location),
		input.Duration,
		input.Distance,
		input.Comment,
		input.Type,
		photoPath,
		&pace,
	)

	if input.Photo != nil {
		err = u.StorageAdapter.UploadMedia(
			ctx,
			configs.MinIoCfg.ProductBucket,
			photoPath,
			input.Photo.File,
			input.Photo.FileSize,
			input.Photo.ContentType,
		)
		if err != nil {
			return err
		}
	}

	err = u.ProductRepository.Create(ctx, *entity)

	if err != nil {
		return err
	}

	entityMarshall, _ := json.Marshal(entity)

	err = u.Messaging.Publish(configs.NatsCfg.ProductTopic, entityMarshall)

	if err != nil {
		log.Println("Error publishing product message:", err)
	}

	return err
}
