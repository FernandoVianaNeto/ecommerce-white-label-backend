package dto

import "mime/multipart"

type PhotoUpload struct {
	File        multipart.File
	FileName    string
	FileSize    int64
	ContentType string
}
type CreateUserInputDto struct {
	Email     string
	BirthDate string
	Name      string
	Password  *string
	Sports    []string
	Origin    string // e.g., "local" or "google"
	Photo     *PhotoUpload
}

type GetUserInputDto struct {
	Uuid string `json:"uuid"`
}

type UpdateUserInputDto struct {
	Uuid      string       `json:"uuid"`
	Email     *string      `json:"email"`
	BirthDate *string      `json:"birth_date"`
	Name      *string      `json:"name"`
	Sports    *[]string    `json:"sports"`
	Photo     *PhotoUpload `json:"photo"`
}

type UserResetPasswordInputDto struct {
	Uuid        string `json:"uuid"`
	NewPassword []byte `json:"new_password"`
}
