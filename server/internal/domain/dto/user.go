package dto

type (
	CreateUser struct {
		Username string `json:"username" validate:"required"`
		password string `json:"password" validate:"required"`
	}
)
