package dto

import "mime/multipart"

type (
	RegisterUser struct {
		Username string `json:"username" validate:"required"`
		Password string `json:"password" validate:"required"`
	}
	DownloadImage struct {
		ImageName string `query:"imageName"`
	}
	UploadImage struct {
		ImageName string                `json:"imageName"`
		File      *multipart.FileHeader `json:"-"`
	}
)
