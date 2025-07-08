package web

import (
	"ecommerce-white-label-backend/internal/domain/dto"
	"ecommerce-white-label-backend/internal/infra/web/requests"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (s *Server) CreateUserHandler(ctx *gin.Context) {
	err := ctx.Request.ParseMultipartForm(10 << 20)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid form data"})
		return
	}

	form := ctx.Request.Form

	password := form.Get("password")

	createUserDto := dto.CreateUserInputDto{
		Email:     form.Get("email"),
		BirthDate: form.Get("birth_date"),
		Name:      form.Get("name"),
		Password:  &password,
		Sports:    form["sports"],
		Origin:    "local",
	}

	fileHeader, err := ctx.FormFile("photo")

	if err == nil {
		file, err := fileHeader.Open()
		createUserDto.Photo = &dto.PhotoUpload{}

		if err == nil {
			createUserDto.Photo.File = file
			createUserDto.Photo.FileSize = fileHeader.Size
			createUserDto.Photo.ContentType = fileHeader.Header.Get("Content-Type")
		}
	}

	err = s.CreateUserUsecase.Execute(ctx, createUserDto)

	if err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}

	ctx.Status(http.StatusOK)
}

func (s *Server) CreateGoogleUserHandler(ctx *gin.Context) {
	var req requests.CreateGoogleUserRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"error": "Invalid request body"})
		return
	}

	err := s.CreateUserUsecase.Execute(ctx, dto.CreateUserInputDto{
		Email:     req.Email,
		BirthDate: req.BirthDate,
		Name:      req.Name,
		Password:  nil,
		Sports:    req.Sports,
		Origin:    "google",
	})

	if err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}

	ctx.Status(http.StatusOK)
}

func (s *Server) GetUserProfileHandler(ctx *gin.Context) {
	var req requests.GetByUuidRequest

	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"error": "Invalid request Uri"})
		return
	}

	response, err := s.GetUserUsecase.Execute(ctx, dto.GetUserInputDto{Uuid: req.Uuid})

	if err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}

	if response == nil {
		ctx.JSON(http.StatusNotFound, "User not found")
		return
	}

	ctx.JSON(http.StatusOK, response)
}

func (s *Server) UpdateUserHandler(ctx *gin.Context) {
	err := ctx.Request.ParseMultipartForm(10 << 20)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid form data"})
		return
	}

	form := ctx.Request.Form

	value := ctx.Value("user_uuid")

	userUuid, ok := value.(string)

	if !ok {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "User uuid not found in context"})
		return
	}

	editUserDto := dto.UpdateUserInputDto{
		Uuid: userUuid,
	}

	if form.Get("birth_date") != "" {
		editUserDto.BirthDate = ptr(form.Get("birth_date"))
	}

	if form.Get("name") != "" {
		editUserDto.Name = ptr(form.Get("name"))
	}

	fmt.Println("DTO", editUserDto)

	fileHeader, err := ctx.FormFile("photo")

	if err == nil {
		file, err := fileHeader.Open()
		editUserDto.Photo = &dto.PhotoUpload{}

		if err == nil {
			editUserDto.Photo.File = file
			editUserDto.Photo.FileSize = fileHeader.Size
			editUserDto.Photo.ContentType = fileHeader.Header.Get("Content-Type")
		}
	}

	err = s.UpdateUserUsecase.Execute(ctx, editUserDto)

	if err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}

	ctx.Status(http.StatusOK)
}

func ptr(s string) *string {
	return &s
}
