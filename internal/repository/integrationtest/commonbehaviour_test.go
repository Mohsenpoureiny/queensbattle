package integrationtest

import (
	"context"
	"fmt"
	"github.com/stretchr/testify/assert"
	"queensbattle/internal/entity"
	"queensbattle/internal/repository"
	"queensbattle/internal/repository/redis"
	"testing"
)

type testType struct {
	ID   string
	Name string
}

func (t testType) EntityID() entity.ID {
	return entity.NewID("testType", t.ID)
}

func TestCommonBehaviourSetAndGet(t *testing.T) {
	redisClient, err := redis.NewRedisClient(fmt.Sprintf("localhost:%s", redisPort))
	defer redisClient.Close()

	assert.NoError(t, err)

	ctx := context.Background()
	cb := repository.NewRedisCommonBehaviour[testType](redisClient)

	// test user 12
	err = cb.Save(ctx, &testType{
		ID:   "12",
		Name: "user12",
	})
	assert.NoError(t, err)

	val, err := cb.Get(ctx, entity.NewID("testType", "12"))

	assert.NoError(t, err)
	assert.Equal(t, "user12", val.Name)
	assert.Equal(t, "12", val.ID)

	// test user 13
	err = cb.Save(ctx, &testType{
		ID:   "13",
		Name: "user13",
	})
	assert.NoError(t, err)

	val, err = cb.Get(ctx, entity.NewID("testType", "13"))

	assert.NoError(t, err)
	assert.Equal(t, "user13", val.Name)
	assert.Equal(t, "13", val.ID)

	// test change id 13 name

	// test user 13
	err = cb.Save(ctx, &testType{
		ID:   "13",
		Name: "changed13",
	})
	assert.NoError(t, err)

	val, err = cb.Get(ctx, entity.NewID("testType", "13"))

	assert.NoError(t, err)
	assert.Equal(t, "changed13", val.Name)

	// not found error

	_, err = cb.Get(ctx, entity.NewID("testType", "14"))
	assert.ErrorIs(t, repository.ErrNotFound, err)
}
