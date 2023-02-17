package domain

import (
	"context"
	"simpleApi/internal/domain/dto"
)

type UserService interface {
	Create(ctx context.Context, p dto.CreateUser) error
}
