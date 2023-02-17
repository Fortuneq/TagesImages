package domain

import (
	"clientTagesImages/internal/domain/dto"
	"context"
)

type UserService interface {
	Create(ctx context.Context, p dto.RegisterUser) error
}
