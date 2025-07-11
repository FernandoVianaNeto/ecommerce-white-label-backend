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
	"fmt"
	"log"
	"strconv"

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
		photoPaths = make([]string, len(input.Photos))
		err        error
	)

	// value := ctx.Value("user_uuid")

	// userUuid, ok := value.(string)

	// if !ok {
	// 	return errors.New("user uuid not found in context")
	// }

	ProductUuid := uuid.New().String()

	if input.Photos != nil {
	}

	for i, photo := range input.Photos {
		photoPath := fmt.Sprintf("/product/%s/%d", ProductUuid, i+1)

		photoPaths[i] = photoPath

		if photo.File != nil {
			err = u.StorageAdapter.UploadMedia(
				ctx,
				configs.MinIoCfg.ProductBucket,
				photoPath,
				photo.File,
				photo.FileSize,
				photo.ContentType,
			)
			if err != nil {
				return err
			}
		}
	}

	priceFloat, err := strconv.ParseFloat(input.Price, 64)
	if err != nil {
		log.Println("Error parsing price:", err)
		return fmt.Errorf("invalid price format: %s", input.Price)
	}

	entity := entity.NewProduct(
		ProductUuid,
		input.Title,
		input.Description,
		priceFloat,
		photoPaths,
		input.Category,
	)

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
