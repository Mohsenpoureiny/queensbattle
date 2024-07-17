package service

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"queensbattle/internal/entity"
	"queensbattle/internal/repository"
	"queensbattle/internal/repository/mocks"
	"testing"
)

func TestAccountService_CreateOrUpdateUserExists(t *testing.T) {
	accRep := &mocks.AccountRepository{}
	s := NewAccountService(accRep)

	accRep.On("Get", mock.Anything, entity.NewID("account", 12)).Return(
		entity.Account{
			ID:        12,
			FirstName: "this_user",
		}, nil).Once()

	accRep.On("Save", mock.Anything, mock.MatchedBy(func(acc entity.Account) bool {
		return acc.FirstName == "user12"
	})).Return(nil).Once()

	newAcc, created, err := s.CreateOrUpdate(context.Background(), entity.Account{
		ID:        12,
		FirstName: "user12",
	})
	assert.NoError(t, err)
	assert.Equal(t, false, created)
	assert.Equal(t, "user12", newAcc.FirstName)

	accRep.AssertExpectations(t)
}

func TestAccountService_CreateOrUpdateUserNotExists(t *testing.T) {
	accRep := &mocks.AccountRepository{}
	s := NewAccountService(accRep)

	accRep.On("Get", mock.Anything, entity.NewID("account", 12)).Return(
		entity.Account{}, repository.ErrNotFound).Once()

	accRep.On("Save", mock.Anything, mock.MatchedBy(func(acc entity.Account) bool {
		return acc.FirstName == "user12"
	})).Return(nil).Once()

	newAcc, created, err := s.CreateOrUpdate(context.Background(), entity.Account{
		ID:        12,
		FirstName: "user12",
	})
	assert.NoError(t, err)
	assert.Equal(t, true, created)
	assert.Equal(t, "user12", newAcc.FirstName)

	accRep.AssertExpectations(t)
}

func TestAccountService_CreateOrUpdateUserHasNotChanged(t *testing.T) {
	accRep := &mocks.AccountRepository{}
	s := NewAccountService(accRep)

	accRep.On("Get", mock.Anything, entity.NewID("account", 12)).Return(
		entity.Account{
			ID:        12,
			FirstName: "user12",
		}, nil).Once()

	newAcc, created, err := s.CreateOrUpdate(context.Background(), entity.Account{
		ID:        12,
		FirstName: "user12",
	})
	assert.NoError(t, err)
	assert.Equal(t, false, created)
	assert.Equal(t, "user12", newAcc.FirstName)

	accRep.AssertExpectations(t)
}
