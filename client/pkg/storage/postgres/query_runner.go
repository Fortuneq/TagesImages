package postgres

import (
	"context"
	"github.com/jmoiron/sqlx"
)

type (
	QueryRunner interface {
		sqlx.ExtContext
		sqlx.Ext
		GetContext(
			ctx context.Context,
			dest interface{},
			query string,
			args ...interface{},
		) error
		SelectContext(
			ctx context.Context,
			dest interface{},
			query string,
			args ...interface{},
		) error
	}
)
