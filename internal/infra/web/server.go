package web

import (
	"context"
	domain_auth_usecase "ecommerce-white-label-backend/internal/domain/usecase/auth"
	domain_product_usecase "ecommerce-white-label-backend/internal/domain/usecase/product"
	domain_usecase "ecommerce-white-label-backend/internal/domain/usecase/user"

	gin "github.com/gin-gonic/gin"
)

type Server struct {
	router            *gin.Engine
	CreateUserUsecase domain_usecase.CreateUserUsecaseInterface
	GetUserUsecase    domain_usecase.GetUserProfileUsecaseInterface
	UpdateUserUsecase domain_usecase.UpdateUserUsecaseInterface
	AuthUseCase       domain_auth_usecase.AuthUsecaseInterface
	GoogleAuthUsecase domain_auth_usecase.GoogleAuthUsecaseInterface

	CreateProductUsecase domain_product_usecase.CreateProductUsecaseInterface
	ListProductsUsecase  domain_product_usecase.ListProductsUsecaseInterface

	// GetProductDetailsUsecase      domain_product_usecase.GetProductDetailsUsecaseInterface
	// AddProductInteractionsUsecase domain_product_usecase.AddInteractionUsecaseInterface
	// GetProductInteractionsUsecase domain_product_usecase.GetInteractionUsecaseInterface
	// DeleteProductUsecase          domain_product_usecase.DeleteProductUsecaseInterface

	GenerateResetPasswordCodeUsecase domain_auth_usecase.GenerateResetPasswordCodeUsecaseInterface
	ResetPasswordUsecase             domain_auth_usecase.ResetPasswordUsecaseInterface
	ValidateResetPasswordCodeUsecase domain_auth_usecase.ValidateResetPasswordCodeUsecaseInterface
}

func NewServer(
	ctx context.Context,
	createUserUsecase domain_usecase.CreateUserUsecaseInterface,
	getUserUsecase domain_usecase.GetUserProfileUsecaseInterface,
	updateUserUsecase domain_usecase.UpdateUserUsecaseInterface,
	authUsecase domain_auth_usecase.AuthUsecaseInterface,
	createProductUsecase domain_product_usecase.CreateProductUsecaseInterface,
	listProductsUsecase domain_product_usecase.ListProductsUsecaseInterface,
	// getProductDetailsUsecase domain_product_usecase.GetProductDetailsUsecaseInterface,
	googleAuthUsecase domain_auth_usecase.GoogleAuthUsecaseInterface,
	generateResetPasswordCodeUsecase domain_auth_usecase.GenerateResetPasswordCodeUsecaseInterface,
	resetPasswordUsecase domain_auth_usecase.ResetPasswordUsecaseInterface,
	validateResetPasswordCodeUsecase domain_auth_usecase.ValidateResetPasswordCodeUsecaseInterface,
	// addProductInteractionUsecase domain_product_usecase.AddInteractionUsecaseInterface,
	// getProductInteractionsUsecase domain_product_usecase.GetInteractionUsecaseInterface,
	// deleteProductUsecase domain_product_usecase.DeleteProductUsecaseInterface,
) *Server {
	router := gin.Default()

	server := &Server{
		CreateUserUsecase:    createUserUsecase,
		GetUserUsecase:       getUserUsecase,
		UpdateUserUsecase:    updateUserUsecase,
		AuthUseCase:          authUsecase,
		CreateProductUsecase: createProductUsecase,
		ListProductsUsecase:  listProductsUsecase,
		// GetProductDetailsUsecase:         getProductDetailsUsecase,
		GoogleAuthUsecase:                googleAuthUsecase,
		GenerateResetPasswordCodeUsecase: generateResetPasswordCodeUsecase,
		ResetPasswordUsecase:             resetPasswordUsecase,
		ValidateResetPasswordCodeUsecase: validateResetPasswordCodeUsecase,
		// AddProductInteractionsUsecase:    addProductInteractionUsecase,
		// GetProductInteractionsUsecase:    getProductInteractionsUsecase,
		// DeleteProductUsecase:             deleteProductUsecase,
	}
	server.router = Routes(router, server)

	return server
}

func (s *Server) Start(address string) error {
	return s.router.Run(address)
}
