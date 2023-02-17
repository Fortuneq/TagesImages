package postgres

//
//import (
//	"context"
//	"fmt"
//	"simpleApi/internal/domain/entity"
//	"simpleApi/internal/domain/repository"
//)
//
//type UserRepository struct {
//	db ch.Conn
//}
//
//var _ repository.User = &UserRepository{}
//
//func (r *UserRepository) Save(ctx context.Context, user entity.User) error {
//	q := "INSERT INTO users.users(username, password) values ($1,$2)"
//	err := r.db.Exec(ctx, q, user.Username, user.Password)
//	if err != nil {
//		return fmt.Errorf("exec: %w", err)
//	}
//	return nil
//}
//
//func NewUserRepository(db ch.Conn) *UserRepository {
//	return &UserRepository{db}
//}
