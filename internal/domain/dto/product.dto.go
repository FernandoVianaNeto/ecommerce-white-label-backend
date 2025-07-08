package dto

type Location struct {
	Lat  string `json:"lat"`
	Long string `json:"long"`
}

type CreateProductInputDto struct {
	Title    string       `json:"title"`
	Location Location     `json:"location"`
	Duration int          `json:"duration"`
	Distance *float64     `json:"distance"`
	Comment  *string      `json:"comment"`
	Type     string       `json:"type"`
	Photo    *PhotoUpload `json:"photo"`
}

type GetProductDetailsInputDto struct {
	Uuid string `json:"uuid"`
}

type ListUserProductsInputDto struct {
	UserUuid string `json:"user_uuid"`
	Page     string `json:"page"`
}

type AddProductInteractionInputDto struct {
	Uuid      string `json:"uuid"`
	UserUuid  string `json:"user_uuid"`
	Emoji     string `json:"emoji"`
	Timestamp string `json:"timestamp"`
}
