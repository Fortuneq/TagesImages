package service

import (
	"context"
	"simpleApi/internal/domain"
	"simpleApi/internal/domain/dto"
	"simpleApi/internal/domain/repository"
	"simpleApi/pkg/govalidator"
)

type User struct {
	userRepo  repository.User
	validator govalidator.Validator
}

func (u *User) Create(ctx context.Context, p dto.CreateUser) error {
	return nil
}

var _ domain.UserService = &User{}

func NewUserService(
	userRepo repository.User,
	validator govalidator.Validator,
) *User {
	return &User{userRepo, validator}
}
