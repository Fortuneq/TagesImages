package service

import (
	"clientTagesImages/internal/domain"
	"clientTagesImages/internal/domain/dto"
	"clientTagesImages/internal/domain/entity"
	"clientTagesImages/internal/domain/repository"
	"clientTagesImages/pkg/govalidator"
	"context"
)

type User struct {
	userRepo  repository.User
	validator govalidator.Validator
}

func (u *User) Create(ctx context.Context, p dto.RegisterUser) error {
	u.userRepo.Save(ctx, entity.User{Password: p.Password, Username: p.Username})
	return nil
}

var _ domain.UserService = &User{}

func NewUserService(
	userRepo repository.User,
	validator govalidator.Validator,
) *User {
	return &User{userRepo, validator}
}
