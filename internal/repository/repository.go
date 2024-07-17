package repository

import (
	"context"
	"errors"
	"queensbattle/internal/entity"
)

var (
	ErrNotFound = errors.New("entity not found")
)

type AccountRepository interface {
	CommonBehaviour[entity.Account]
}

type CommonBehaviour[T entity.Entity] interface {
	Get(ctx context.Context, id entity.ID) (T, error)
	Save(ctx context.Context, ent entity.Entity) error
}
