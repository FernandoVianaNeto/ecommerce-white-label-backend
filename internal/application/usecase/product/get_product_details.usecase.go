package product_usecase

// import (
// 	"context"
// 	configs "ecommerce-white-label-backend/cmd/config"
// 	storage_adapter "ecommerce-white-label-backend/internal/domain/adapters/storage"
// 	"ecommerce-white-label-backend/internal/domain/dto"
// 	"ecommerce-white-label-backend/internal/domain/entity"
// 	domain_repository "ecommerce-white-label-backend/internal/domain/repository"
// 	domain_product_usecase "ecommerce-white-label-backend/internal/domain/usecase/product"
// 	"strconv"
// 	"time"
// )

// type GetProductDetailsUsecase struct {
// 	ProductRepository domain_repository.ProductRepositoryInterface
// 	StorageAdapter    storage_adapter.StorageAdapterInterface
// }

// func NewGetProductDetailsUseCase(
// 	repository domain_repository.ProductRepositoryInterface,
// 	storageAdapter storage_adapter.StorageAdapterInterface,
// ) domain_product_usecase.GetProductDetailsUsecaseInterface {
// 	return &GetProductDetailsUsecase{
// 		ProductRepository: repository,
// 		StorageAdapter:    storageAdapter,
// 	}
// }

// func (u *GetProductDetailsUsecase) Execute(ctx context.Context, input dto.GetProductDetailsInputDto) (*entity.Product, error) {
// 	repositoryResponse, err := u.ProductRepository.GetByUuid(ctx, input.Uuid)

// 	if err != nil {
// 		if err.Error() == "mongo: no documents in result" {
// 			return nil, nil
// 		}
// 	}

// 	urlExpiration, err := strconv.Atoi(configs.MinIoCfg.PresignedURLExpiration)

// 	if err != nil {
// 		return nil, err
// 	}

// 	repositoryResponse.Photo, err = u.StorageAdapter.GeneratePresignedURL(
// 		ctx,
// 		configs.MinIoCfg.ProductBucket,
// 		repositoryResponse.Photo,
// 		time.Duration(urlExpiration)*time.Minute,
// 	)

// 	if err != nil {
// 		return nil, err
// 	}

// 	return repositoryResponse, err
// }
