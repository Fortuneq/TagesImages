package repository

import (
	"clientTagesImages/internal/domain/entity"
	"context"
	"github.com/jmoiron/sqlx"
)

type User interface {
	Save(ctx context.Context, user entity.User) error
}

func (u *UserRepo) Save(ctx context.Context, user entity.User) error {
	u.db.ExecContext(ctx, `INSERT INTO clients values($1,$2)`, user.Username, user.Password)
	return nil
}

type UserRepo struct {
	db *sqlx.DB
}

func NewUserRepo(
	db *sqlx.DB,
) *UserRepo {
	return &UserRepo{db: db}
}

var _ User = &UserRepo{}
