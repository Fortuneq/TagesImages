package repository

import (
	"context"
	"simpleApi/internal/domain/entity"
)

type User interface {
	Save(ctx context.Context, user entity.User) error
}
