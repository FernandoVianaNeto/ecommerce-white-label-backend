package domain_response

type GetUserProfileResponse struct {
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Sports    *[]string `json:"sports"`
	BirthDate string    `json:"birth_date"`
	PhotoUrl  string    `json:"photo_url"`
}
