package service

import (
	"context"
	"errors"
	"queensbattle/internal/entity"
	"queensbattle/internal/repository"
	"time"
)

const (
	DefaultState = "home"
)

type AccountService struct {
	accounts repository.AccountRepository
}

func NewAccountService(rep repository.AccountRepository) *AccountService {
	return &AccountService{accounts: rep}
}

func (a *AccountService) Update(ctx context.Context, account entity.Account) error {
	return a.accounts.Save(ctx, account)
}

func (a *AccountService) CreateOrUpdate(ctx context.Context, account entity.Account) (entity.Account, bool, error) {
	savedAccount, err := a.accounts.Get(ctx, account.EntityID())

	// user exists in the database
	if err == nil {
		if savedAccount.Username != account.Username || savedAccount.FirstName != account.FirstName {
			savedAccount.Username = account.Username
			savedAccount.FirstName = account.FirstName
			return savedAccount, false, a.accounts.Save(ctx, savedAccount)
		}
		return savedAccount, false, nil
	}
	// user does not exists in the database

	if errors.Is(err, repository.ErrNotFound) {
		account.JoinedAt = time.Now()
		account.State = DefaultState
		return account, true, a.accounts.Save(ctx, account)
	}

	return entity.Account{}, false, err
}
