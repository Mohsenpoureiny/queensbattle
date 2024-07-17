package repository

import (
	"context"
	"errors"
	"queensbattle/internal/entity"
	"queensbattle/pkg/jsonhelper"

	"github.com/redis/rueidis"
	"github.com/sirupsen/logrus"
)

var _ CommonBehaviou[entity.Entity] = &RedisCommandBehaviour[entity.Entity]{}

type CommonBehaviou[T entity.Entity] interface {
	Get(ctx context.Context, id entity.ID) (T, error)
	Save(ctx context.Context, ent entity.Entity) error
}

type RedisCommandBehaviour[T entity.Entity] struct {
	client rueidis.Client
}

func NewRedisCommonBehaviour[T entity.Entity](client rueidis.Client) RedisCommandBehaviour[T] {
	return RedisCommandBehaviour[T]{
		client: client,
	}
}
func (r RedisCommandBehaviour[T]) Get(ctx context.Context, id entity.ID) (T, error) {
	var t T
	cmd := r.client.B().JsonGet().Key(id.String()).Path(".").Build()
	val, err := r.client.Do(ctx, cmd).ToString()
	if err != nil {
		// handle redis nil error
		if errors.Is(err, rueidis.Nil) {
			return t, ErrNotFound
		}
		logrus.WithError(err).WithField("id", id).Errorln("couldn't retrive")
		return t, err
	}
	return jsonhelper.Decode[T]([]byte(val)), nil
}

func (r RedisCommandBehaviour[T]) Save(ctx context.Context, ent entity.Entity) error {
	cmd := r.client.B().JsonSet().Key(ent.EntityID().String()).Path("$").Value(string(jsonhelper.Encode(ent))).Build()

	if err := r.client.Do(ctx, cmd).Error(); err != nil {
		logrus.WithError(err).WithField("ent", ent).Errorln("couldn't save the entity")
		return err
	}
	return nil
}
