package requests

type CreateGoogleUserRequest struct {
	Email     string   `json:"email"`
	BirthDate string   `json:"birth_date"`
	Name      string   `json:"name"`
	Password  string   `json:"password"`
	Sports    []string `json:"sports"`
	GoogleSub string   `json:"google_sub"`
}

type GetByUuidRequest struct {
	Uuid string `uri:"uuid"`
}
