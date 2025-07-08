package web

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (s *Server) CreateProductHandler(ctx *gin.Context) {
	// err := ctx.Request.ParseMultipartForm(10 << 20)

	// if err != nil {
	// 	ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid form data"})
	// 	return
	// }

	// form := ctx.Request.Form

	// if form.Get("title") == "" ||
	// 	form.Get("loc_latitude") == "" ||
	// 	form.Get("loc_longitude") == "" ||
	// 	form.Get("duration") == "" ||
	// 	form.Get("type") == "" {
	// 	ctx.JSON(http.StatusBadRequest, gin.H{"error": "Missing required fields"})
	// 	return
	// }

	// durationInt, err := strconv.Atoi(form.Get("duration"))
	// if err != nil {
	// 	ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid duration"})
	// 	return
	// }

	// var distanceFloat float64
	// if form.Get("duration") != "" {
	// 	distanceFloat, err = strconv.ParseFloat(form.Get("duration"), 64)
	// 	if err != nil {
	// 		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid duration"})
	// 		return
	// 	}
	// }

	// var comment string
	// if form.Get("comment") != "" {
	// 	distanceFloat, err = strconv.ParseFloat(form.Get("duration"), 64)
	// 	if err != nil {
	// 		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid comment"})
	// 		return
	// 	}
	// }

	// createProductDto := dto.CreateProductInputDto{
	// 	Title: form.Get("title"),
	// 	Location: dto.Location{
	// 		Lat:  form.Get("loc_latitude"),
	// 		Long: form.Get("loc_longitude"),
	// 	},
	// 	Duration: durationInt,
	// 	Distance: &distanceFloat,
	// 	Comment:  &comment,
	// 	Type:     form.Get("type"),
	// }

	// fileHeader, err := ctx.FormFile("photo")

	// if err == nil {
	// 	file, err := fileHeader.Open()
	// 	createProductDto.Photo = &dto.PhotoUpload{}

	// 	if err == nil {
	// 		createProductDto.Photo.File = file
	// 		createProductDto.Photo.FileSize = fileHeader.Size
	// 		createProductDto.Photo.ContentType = fileHeader.Header.Get("Content-Type")
	// 	}
	// }

	// err = s.CreateProductUsecase.Execute(ctx, createProductDto)

	// if err != nil {
	// 	ctx.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
	// 	return
	// }

	ctx.Status(http.StatusOK)
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

// func (s *Server) ListUserProducts(ctx *gin.Context) {
// 	var queryParams requests.ListUserProductsQueryParams
// 	var req requests.ListUserProductsRequest

// 	if err := ctx.ShouldBindUri(&req); err != nil {
// 		ctx.JSON(http.StatusBadRequest, "Invalid param")
// 		return
// 	}

// 	response, err := s.ListUserProductsUsecase.Execute(ctx, dto.ListUserProductsInputDto{UserUuid: req.UserUuid, Page: queryParams.Page})

// 	if err != nil {
// 		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
// 		return
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
