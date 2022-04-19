package mocks

import (
	"coinwallet/wallet"
	"github.com/google/uuid"
)

type RepositoryMock struct {
}

func (mock RepositoryMock) Save(w wallet.Wallet) {
}

func (mock RepositoryMock) Update(w wallet.Wallet) wallet.Wallet {
	return w
}

func (mock RepositoryMock) FindAll() []wallet.Wallet {
	return nil
}

var FindByIdMock func(id uuid.UUID) wallet.Wallet

func (mock RepositoryMock) FindById(id uuid.UUID) wallet.Wallet {
	return FindByIdMock(id)
}
