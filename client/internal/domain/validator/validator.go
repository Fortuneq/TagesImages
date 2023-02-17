package validator

import "context"

type Validator interface {
	Validate(ctx context.Context, v any) error
}
