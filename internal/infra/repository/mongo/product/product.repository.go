package product_mongo_repository

import (
	"context"
	configs "ecommerce-white-label-backend/cmd/config"
	"ecommerce-white-label-backend/internal/domain/dto"
	"ecommerce-white-label-backend/internal/domain/entity"
	domain_repository "ecommerce-white-label-backend/internal/domain/repository"
	domain_response "ecommerce-white-label-backend/internal/domain/response"
	"errors"
	"fmt"
	"strconv"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type ProductRepository struct {
	db         *mongo.Database
	collection *mongo.Collection
}

func NewProductRepository(db *mongo.Database) domain_repository.ProductRepositoryInterface {
	collection := db.Collection(configs.MongoCfg.ProductCollection)

	return &ProductRepository{
		db:         db,
		collection: collection,
	}
}

func (f *ProductRepository) Create(ctx context.Context, input entity.Product) error {
	_, err := f.collection.InsertOne(ctx, ProductModel{
		Uuid:      input.Uuid,
		UserUuid:  input.UserUuid,
		Title:     input.Title,
		Location:  Location(input.Location),
		Duration:  input.Duration,
		Pace:      input.Pace,
		Distance:  input.Distance,
		Comment:   input.Comment,
		Type:      input.Type,
		Photo:     input.Photo,
		CreatedAt: input.CreatedAt,
		UpdatedAt: input.UpdatedAt,
	})

	return err
}

func (f *ProductRepository) GetByUuid(ctx context.Context, uuid string) (*entity.Product, error) {
	var model ProductModel

	filter := bson.M{
		"uuid": uuid,
	}

	err := f.collection.FindOne(ctx, filter).Decode(&model)

	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, nil
		}

		return nil, err
	}

	reactionSummary := make(map[string]int)
	for _, r := range model.Reactions {
		reactionSummary[r.Emoji]++
	}

	entity := entity.Product{
		Uuid:            model.Uuid,
		UserUuid:        model.UserUuid,
		Title:           model.Title,
		Location:        entity.Location(model.Location),
		Duration:        model.Duration,
		Pace:            model.Pace,
		Distance:        model.Distance,
		Comment:         model.Comment,
		Type:            model.Type,
		Photo:           model.Photo,
		CreatedAt:       model.CreatedAt,
		UpdatedAt:       model.UpdatedAt,
		ReactionSummary: reactionSummary,
		Reactions:       ConvertReactions(model.Reactions),
	}
	// TODO REFATORAR USE CASE DO GET Product DETAILS PARA RETORNAR APENAS O NECESSÁRIO

	return &entity, err
}

func (f *ProductRepository) ListUserProducts(ctx context.Context, userUuid string, pageStr string) (domain_response.ListProductsPaginatedResponse, error) {
	page, err := strconv.Atoi(pageStr)
	if err != nil || page < 1 {
		page = 1
	}

	defaultMetadata := domain_response.GetMetadataParams(page, 10)
	defaultResponse := domain_response.ListProductsPaginatedResponse{
		Items:    []entity.Product{},
		Metadata: defaultMetadata,
	}

	limit := int64(domain_response.DEFAULT_ITEMS_PER_PAGE)
	skip := int64((page - 1)) * limit

	filter := bson.M{
		"user_uuid": userUuid,
	}

	opts := options.Find()
	opts.SetLimit(limit)
	opts.SetSkip(skip)
	opts.SetSort(bson.M{"created_at": -1})

	total, err := f.collection.CountDocuments(ctx, filter)

	if err != nil {
		return defaultResponse, err
	}

	cursor, err := f.collection.Find(ctx, filter, opts)
	if err != nil {
		return defaultResponse, err
	}

	var Products []ProductModel
	if err = cursor.All(ctx, &Products); err != nil {
		return defaultResponse, err
	}

	entitiesProduct := make([]entity.Product, 0, len(Products))
	for _, Product := range Products {
		reactionSummary := make(map[string]int)
		for _, r := range Product.Reactions {
			reactionSummary[r.Emoji]++
		}

		entitiesProduct = append(entitiesProduct, entity.Product{
			Uuid:            Product.Uuid,
			UserUuid:        Product.UserUuid,
			Title:           Product.Title,
			Location:        entity.Location(Product.Location),
			Duration:        Product.Duration,
			Pace:            Product.Pace,
			Distance:        Product.Distance,
			Comment:         Product.Comment,
			Type:            Product.Type,
			Photo:           Product.Photo,
			CreatedAt:       Product.CreatedAt,
			UpdatedAt:       Product.UpdatedAt,
			ReactionSummary: reactionSummary,
		})
	}

	response := domain_response.ListProductsPaginatedResponse{
		Items:    entitiesProduct,
		Metadata: domain_response.GetMetadataParams(page, total),
	}

	return response, nil
}

func (f *ProductRepository) AddInteraction(ctx context.Context, input dto.AddProductInteractionInputDto) error {
	filter := bson.M{
		"uuid": input.Uuid,
	}

	interaction := entity.EmojiReaction{
		UserUuid:  input.UserUuid,
		Emoji:     input.Emoji,
		Timestamp: input.Timestamp,
	}

	update := bson.M{
		"$push": bson.M{
			"reactions": interaction,
		},
	}

	opts := options.Update().SetUpsert(true)

	_, err := f.collection.UpdateOne(ctx, filter, update, opts)
	if err != nil {
		return err
	}

	return nil
}

func (f *ProductRepository) Delete(ctx context.Context, ProductUuid string) error {
	filter := bson.M{
		"uuid": ProductUuid,
	}

	fmt.Println(filter)

	_, err := f.collection.DeleteOne(ctx, filter)
	if err != nil {
		return err
	}

	return nil
}

func ConvertReactions(models []EmojiReaction) []entity.EmojiReaction {
	reactions := make([]entity.EmojiReaction, 0, len(models))
	for _, m := range models {
		reactions = append(reactions, entity.EmojiReaction{
			UserUuid:  m.UserUuid,
			Emoji:     m.Emoji,
			Timestamp: m.Timestamp,
		})
	}
	return reactions
}
