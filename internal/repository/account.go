package repository

import (
	"github.com/redis/rueidis"
	"queensbattle/internal/entity"
)

var _ AccountRepository = &AccountRedisRepository{}

type AccountRedisRepository struct {
	*RedisCommandBehaviour[entity.Account]
}

func NewAccountRedisRepository(client rueidis.Client) *AccountRedisRepository {
	return &AccountRedisRepository{
		NewRedisCommonBehaviour[entity.Account](client),
	}
}
