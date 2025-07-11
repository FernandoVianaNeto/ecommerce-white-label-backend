package product_mongo_repository

type ProductModel struct {
	Uuid        string   `bson:"uuid" json:"uuid"`
	Title       string   `bson:"title" json:"title"`
	Description string   `bson:"description" json:"description"`
	Price       float64  `bson:"price" json:"price"`
	Category    string   `bson:"category" json:"category"`
	Photos      []string `bson:"photos" json:"photos"`
	CreatedAt   string   `bson:"created_at" json:"created_at"`
	UpdatedAt   string   `bson:"updated_at" json:"updated_at"`
}
