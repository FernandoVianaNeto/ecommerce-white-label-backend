package product_mongo_repository

type Location struct {
	Lat  string `json:"lat"`
	Long string `json:"long"`
}

type EmojiReaction struct {
	UserUuid  string `json:"user_uuid" bson:"user_uuid"`
	Emoji     string `json:"emoji" bson:"emoji"`
	Timestamp string `json:"timestamp" bson:"timestamp"`
}

type ProductModel struct {
	Uuid      string          `bson:"uuid" json:"uuid"`
	UserUuid  string          `bson:"user_uuid" json:"user_uuid"`
	Title     string          `bson:"title" json:"title"`
	Location  Location        `bson:"location" json:"location"`
	Duration  int             `bson:"duration" json:"duration"`
	Pace      *string         `bson:"pace" json:"pace"`
	Distance  *float64        `bson:"distance" json:"distance"`
	Comment   *string         `bson:"comment" json:"comment"`
	Type      string          `bson:"type" json:"type"`
	Photo     string          `bson:"photo" json:"photo"`
	CreatedAt string          `bson:"created_at" json:"created_at"`
	UpdatedAt string          `bson:"updated_at" json:"updated_at"`
	Reactions []EmojiReaction `json:"reactions,omitempty" bson:"reactions,omitempty"`
}
