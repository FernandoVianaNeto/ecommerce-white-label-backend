package mongo_repository

type UserModel struct {
	Uuid         string    `bson:"uuid" json:"uuid"`
	Email        string    `bson:"email" json:"email"`
	Name         string    `bson:"name" json:"name"`
	BirthDate    string    `bson:"birth_date" json:"birth_date"`
	Password     *string   `bson:"password,omitempty" json:"password,omitempty"`
	Sports       *[]string `bson:"sports,omitempty" json:"sports,omitempty"`
	AuthProvider string    `bson:"auth_provider" json:"auth_provider"`
	Photo        string    `bson:"photo" json:"photo"`
}
