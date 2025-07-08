package domain_response

import "ecommerce-white-label-backend/internal/domain/entity"

type ListProductsPaginatedResponse struct {
	Items    []entity.Product   `json:"items"`
	Metadata PaginationMetadata `json:"metadata"`
}

type ProductInteractionResponse struct {
	UserUuid string `json:"user_uuid"`
	Photo    string `json:"photo"`
	Name     string `json:"name"`
	Emoji    string `json:"emoji"`
}
