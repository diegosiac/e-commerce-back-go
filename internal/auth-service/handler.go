package authservice

import (
	"context"
)

type Resolver struct{}

func (r *Resolver) Hello(ctx context.Context) (string, error) {
	return "Hello, world!", nil
}
