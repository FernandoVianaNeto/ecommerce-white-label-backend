package app

import (
	"context"

	configs "ecommerce-white-label-backend/cmd/config"
	service "ecommerce-white-label-backend/internal/application/services"
	auth_usecase "ecommerce-white-label-backend/internal/application/usecase/auth"
	product_usecase "ecommerce-white-label-backend/internal/application/usecase/product"
	user_usecase "ecommerce-white-label-backend/internal/application/usecase/users"
	adapter "ecommerce-white-label-backend/internal/domain/adapters/email_sender"
	"ecommerce-white-label-backend/internal/domain/adapters/messaging"
	storage_adapter "ecommerce-white-label-backend/internal/domain/adapters/storage"
	domain_repository "ecommerce-white-label-backend/internal/domain/repository"
	domain_service "ecommerce-white-label-backend/internal/domain/service"
	domain_auth_usecase "ecommerce-white-label-backend/internal/domain/usecase/auth"
	domain_product_usecase "ecommerce-white-label-backend/internal/domain/usecase/product"
	domain_usecase "ecommerce-white-label-backend/internal/domain/usecase/user"
	"ecommerce-white-label-backend/internal/infra/adapter/minio"
	"ecommerce-white-label-backend/internal/infra/adapter/sendgrid"
	Product_mongo_repository "ecommerce-white-label-backend/internal/infra/repository/mongo/product"
	reset_password_code_mongo_repository "ecommerce-white-label-backend/internal/infra/repository/mongo/reset_password_code"
	mongo_repository "ecommerce-white-label-backend/internal/infra/repository/mongo/user"
	"ecommerce-white-label-backend/internal/infra/web"
	mongoPkg "ecommerce-white-label-backend/pkg/mongo"
	natsclient "ecommerce-white-label-backend/pkg/nats"
	"ecommerce-white-label-backend/pkg/storage"

	"go.mongodb.org/mongo-driver/mongo"
)

type Application struct {
	UseCases UseCases
}
type UseCases struct {
	userUseCase                      domain_usecase.CreateUserUsecaseInterface
	GetUserUsecase                   domain_usecase.GetUserProfileUsecaseInterface
	UpdateUserUsecase                domain_usecase.UpdateUserUsecaseInterface
	AuthUsecase                      domain_auth_usecase.AuthUsecaseInterface
	GoogleAuthUsecase                domain_auth_usecase.GoogleAuthUsecaseInterface
	CreateProductUsecase             domain_product_usecase.CreateProductUsecaseInterface
	ListProductsUsecase              domain_product_usecase.ListProductsUsecaseInterface
	GetProductDetailsUsecase         domain_product_usecase.GetProductDetailsUsecaseInterface
	AddProductInteractionsUsecase    domain_product_usecase.AddInteractionUsecaseInterface
	GetProductInteractionUsecase     domain_product_usecase.GetInteractionUsecaseInterface
	DeleteProductUsecase             domain_product_usecase.DeleteProductUsecaseInterface
	GenerateResetPasswordCodeUsecase domain_auth_usecase.GenerateResetPasswordCodeUsecaseInterface
	ResetPasswordUsecase             domain_auth_usecase.ResetPasswordUsecaseInterface
	ValidateResetPasswordCodeUsecase domain_auth_usecase.ValidateResetPasswordCodeUsecaseInterface
}

type Services struct {
	encryptStringService domain_service.EncryptStringServiceInterface
}

type Adapters struct {
	emailSenderAdapter adapter.EmailSenderAdapterInterface
	storageAdapter     storage_adapter.StorageAdapterInterface
}

type Repositories struct {
	UserRepository              domain_repository.UserRepositoryInterface
	ProductRepository           domain_repository.ProductRepositoryInterface
	ResetPasswordCodeRepository domain_repository.ResetPasswordCodeRepositoryInterface
}

func NewApplication() *web.Server {
	ctx := context.Background()

	mongoConnectionInput := mongoPkg.MongoInput{
		DSN:      configs.MongoCfg.Dsn,
		Database: configs.MongoCfg.Database,
	}

	db := mongoPkg.NewMongoDatabase(ctx, mongoConnectionInput)

	eventClient := natsclient.New(configs.NatsCfg.Host)
	eventClient.Connect()

	repositories := NewRepositories(ctx, db)

	services := NewServices(ctx)

	adapters := NewAdapters(ctx)

	usecases := NewUseCases(
		ctx,
		repositories.UserRepository,
		repositories.ProductRepository,
		repositories.ResetPasswordCodeRepository,
		services,
		adapters,
		eventClient,
	)

	srv := web.NewServer(
		ctx,
		usecases.userUseCase,
		usecases.GetUserUsecase,
		usecases.UpdateUserUsecase,
		usecases.AuthUsecase,
		usecases.CreateProductUsecase,
		usecases.ListProductsUsecase,
		// usecases.GetProductDetailsUsecase,
		usecases.GoogleAuthUsecase,
		usecases.GenerateResetPasswordCodeUsecase,
		usecases.ResetPasswordUsecase,
		usecases.ValidateResetPasswordCodeUsecase,
		// usecases.AddProductInteractionsUsecase,
		// usecases.GetProductInteractionUsecase,
		// usecases.DeleteProductUsecase,
	)

	return srv
}

func NewRepositories(
	ctx context.Context,
	db *mongo.Database,
) Repositories {
	userRepository := mongo_repository.NewUserRepository(db)
	ProductRepository := Product_mongo_repository.NewProductRepository(db)
	resetPasswordCodeRepository := reset_password_code_mongo_repository.NewResetPasswordCodeRepository(db)

	return Repositories{
		UserRepository:              userRepository,
		ProductRepository:           ProductRepository,
		ResetPasswordCodeRepository: resetPasswordCodeRepository,
	}
}

func NewServices(
	ctx context.Context,
) Services {
	encryptStringService := service.NewEncryptStringService()

	return Services{
		encryptStringService: encryptStringService,
	}
}

func NewAdapters(
	ctx context.Context,
) Adapters {
	emailSenderAdapter := sendgrid.NewEmailSenderAdapter(ctx)
	minioAdapter := NewStorageAdapter(ctx)

	return Adapters{
		emailSenderAdapter: emailSenderAdapter,
		storageAdapter:     minioAdapter,
	}
}

func NewUseCases(
	ctx context.Context,
	userRepository domain_repository.UserRepositoryInterface,
	ProductRepository domain_repository.ProductRepositoryInterface,
	resetPasswordCodeRepository domain_repository.ResetPasswordCodeRepositoryInterface,
	services Services,
	adapters Adapters,
	eventClient messaging.Client,
) UseCases {
	userUsecase := user_usecase.NewCreateUserUseCase(userRepository, services.encryptStringService, adapters.storageAdapter)
	getUserUsecase := user_usecase.NewGetUserProfileUseCase(userRepository, adapters.storageAdapter)
	updateUserUsecase := user_usecase.NewUpdateUserUseCase(userRepository, adapters.storageAdapter)

	//AUTH
	authUsecase := auth_usecase.NewAuthUsecase(userRepository)
	googleAuthUsecase := auth_usecase.NewGoogleAuthUsecase(userRepository)
	generateResetPasswordCodeUsecase := auth_usecase.NewGenerateResetPasswordCodeUsecase(resetPasswordCodeRepository, userRepository, adapters.emailSenderAdapter)
	resetPasswordUsecase := auth_usecase.NewResetPasswordUsecase(userRepository, resetPasswordCodeRepository, services.encryptStringService)
	validateResetPasswordCodeUsecase := auth_usecase.NewValidateResetPasswordCodeUsecase(resetPasswordCodeRepository)

	createProductUsecase := product_usecase.NewCreateProductUseCase(ProductRepository, adapters.storageAdapter, eventClient)
	listProductsUsecase := product_usecase.NewListProductsUseCase(ProductRepository)
	// getProductDetailsUsecase := product_usecase.NewGetProductDetailsUseCase(ProductRepository, adapters.storageAdapter)
	// addProductInteractionUsecase := product_usecase.NewAddInteractionUseCase(ProductRepository)
	// getInteractionsUsecase := product_usecase.NewGetProductInteractionUsecase(ProductRepository, userRepository, adapters.storageAdapter)
	deleteProductUsecase := product_usecase.NewDeleteProductUseCase(ProductRepository)

	return UseCases{
		userUseCase:          userUsecase,
		GetUserUsecase:       getUserUsecase,
		UpdateUserUsecase:    updateUserUsecase,
		AuthUsecase:          authUsecase,
		GoogleAuthUsecase:    googleAuthUsecase,
		CreateProductUsecase: createProductUsecase,
		ListProductsUsecase:  listProductsUsecase,
		// GetProductDetailsUsecase:         getProductDetailsUsecase,
		GenerateResetPasswordCodeUsecase: generateResetPasswordCodeUsecase,
		ResetPasswordUsecase:             resetPasswordUsecase,
		ValidateResetPasswordCodeUsecase: validateResetPasswordCodeUsecase,
		// AddProductInteractionsUsecase:    addProductInteractionUsecase,
		// GetProductInteractionUsecase:     getInteractionsUsecase,
		DeleteProductUsecase: deleteProductUsecase,
	}
}

func NewStorageAdapter(
	ctx context.Context,
) storage_adapter.StorageAdapterInterface {
	client, err := storage.NewMinioClient(
		configs.MinIoCfg.Host,
		configs.MinIoCfg.User,
		configs.MinIoCfg.Password,
	)

	if err != nil {
		panic("Failed to create MinIO client: " + err.Error())
	}

	err = storage.CreateBucketIfNotExists(ctx, client, configs.MinIoCfg.ProfileBucket)

	if err != nil {
		panic("Failed to create profile bucket: " + err.Error())
	}

	err = storage.CreateBucketIfNotExists(ctx, client, configs.MinIoCfg.ProductBucket)

	if err != nil {
		panic("Failed to create product bucket: " + err.Error())
	}

	minioAdapter := minio.NewMinIoAdapter(ctx, client)

	return minioAdapter
}
