package biz

import (
	"context"
)

// Greeter is a Greeter model.
type Greeter struct {
	Hello string
}

// GreeterRepo is a Greater repo.
type GreeterRepo interface {
	NewID(ctx context.Context) (int64, error)
}

// GreeterUseCase is a Greeter use case.
type GreeterUseCase struct {
	repo GreeterRepo
}

// NewGreeterUseCase new a Greeter use case.
func NewGreeterUseCase(repo GreeterRepo) *GreeterUseCase {
	return &GreeterUseCase{repo: repo}
}

func (g *GreeterUseCase) NewID(ctx context.Context) (int64, error) {
	return g.repo.NewID(ctx)
}
