package web

import (
	"ecommerce-white-label-backend/internal/domain/dto"
	"ecommerce-white-label-backend/internal/infra/web/requests"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (s *Server) CreateProductHandler(ctx *gin.Context) {
	err := ctx.Request.ParseMultipartForm(10 << 20)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid form data"})
		return
	}

	form := ctx.Request.Form

	if form.Get("title") == "" ||
		form.Get("description") == "" ||
		form.Get("category") == "" ||
		form.Get("price") == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Missing required fields"})
		return
	}

	createProductDto := dto.CreateProductInputDto{
		Title:       form.Get("title"),
		Description: form.Get("description"),
		Price:       form.Get("price"),
		Category:    form.Get("category"),
		Photos:      []*dto.PhotoUpload{},
	}

	multipartForm, err := ctx.MultipartForm()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Erro ao ler arquivos"})
		return
	}

	files := multipartForm.File["photos"]
	if len(files) == 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Envie pelo menos uma foto"})
		return
	}

	for _, fileHeader := range files {
		file, err := fileHeader.Open()
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao abrir arquivo"})
			return
		}

		photo := &dto.PhotoUpload{
			File:        file,
			FileSize:    fileHeader.Size,
			ContentType: fileHeader.Header.Get("Content-Type"),
			FileName:    fileHeader.Filename,
		}

		createProductDto.Photos = append(createProductDto.Photos, photo)
	}

	err = s.CreateProductUsecase.Execute(ctx, createProductDto)

	if err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}

	ctx.Status(http.StatusOK)
}

func (s *Server) ListProducts(ctx *gin.Context) {
	var queryParams requests.ListUserProductsQueryParams

	if err := ctx.ShouldBindQuery(&queryParams); err != nil {
		ctx.JSON(http.StatusBadRequest, "Invalid query")
		return
	}

	response, err := s.ListProductsUsecase.Execute(ctx, queryParams.Page)

	if err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, response)
}

// func (s *Server) GetProductDetails(ctx *gin.Context) {
// 	var req requests.GetProductDetailsRequest

// 	if err := ctx.ShouldBindUri(&req); err != nil {
// 		ctx.JSON(http.StatusBadRequest, "Invalid param")
// 		return
// 	}

// 	response, err := s.GetProductDetailsUsecase.Execute(ctx, dto.GetProductDetailsInputDto{Uuid: req.Uuid})

// 	if err != nil {
// 		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
// 		return
// 	}

// 	if response == nil {
// 		ctx.JSON(http.StatusNotFound, "Product not found")
// 	}

// 	ctx.JSON(http.StatusOK, response)
// }

// func (s *Server) AddInteraction(ctx *gin.Context) {
// 	var req requests.AddProductInteractionRequest

// 	if err := ctx.ShouldBindJSON(&req); err != nil {
// 		ctx.JSON(http.StatusBadRequest, "Invalid Body")
// 		return
// 	}

// 	err := s.AddProductInteractionsUsecase.Execute(ctx, dto.AddProductInteractionInputDto{
// 		Uuid:  req.Uuid,
// 		Emoji: req.Emoji,
// 	})

// 	if err != nil {
// 		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
// 		return
// 	}

// 	ctx.Status(http.StatusCreated)
// }

// func (s *Server) GetInteractions(ctx *gin.Context) {
// 	var req requests.GetProductInteractionRequest

// 	if err := ctx.ShouldBindUri(&req); err != nil {
// 		ctx.JSON(http.StatusBadRequest, "Invalid Uuid param")
// 		return
// 	}

// 	response, err := s.GetProductInteractionsUsecase.Execute(ctx, req.Uuid)

// 	if err != nil {
// 		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
// 		return
// 	}

// 	ctx.JSON(http.StatusOK, response)
// }

// func (s *Server) DeleteProductHandler(ctx *gin.Context) {
// 	var req requests.GetProductInteractionRequest

// 	if err := ctx.ShouldBindUri(&req); err != nil {
// 		ctx.JSON(http.StatusBadRequest, "Invalid Uuid param")
// 		return
// 	}

// 	err := s.DeleteProductUsecase.Execute(ctx, req.Uuid)

// 	if err != nil {
// 		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
// 		return
// 	}

// 	ctx.Status(http.StatusOK)
// }
