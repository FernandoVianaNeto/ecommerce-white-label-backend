package web

import (
	"ecommerce-white-label-backend/internal/infra/web/middleware"

	"github.com/gin-gonic/gin"
)

func Routes(engine *gin.Engine, server *Server) *gin.Engine {
	{
		auth := engine.Group("/v1/auth")
		{
			auth.POST("/", server.AuthHandler)
			auth.POST("/google", server.GoogleAuthHandler)
			auth.POST("/generate-reset-code", server.GenerateResetPasswordCodeHandler)
			auth.POST("/reset-password", server.ResetPasswordHandler)
			auth.POST("/validate-code", server.ValidateResetPasswordCode)
		}
	}

	{
		user := engine.Group("/v1/user")
		{
			user.POST("/create", server.CreateUserHandler)
			user.POST("/create/google", server.CreateGoogleUserHandler)
			user.GET("/:uuid/profile", middleware.JwtAuthMiddleware(), server.GetUserProfileHandler)
			user.PUT("/", middleware.JwtAuthMiddleware(), server.UpdateUserHandler)
		}
	}

	{
		product := engine.Group("/product")
		// product := engine.Group("/v1/product", middleware.JwtAuthMiddleware())
		{
			product.POST("/", server.CreateProductHandler)
			product.GET("/list", server.ListProducts)
			// product.GET("/:uuid", server.GetproductDetails)
			// product.GET("/user/:uuid", server.ListUserproducts)
			// product.GET("/:uuid/interactions", server.GetInteractions)
			// product.DELETE("/:uuid", server.DeleteWorkoutHandler)
		}
	}

	{
		heathCheck := engine.Group("/v1/health")
		{
			heathCheck.GET("/check", server.HealthCheckHandler)
		}
	}

	return engine
}
