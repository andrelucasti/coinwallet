package wallet

import (
	"time"

	"github.com/google/uuid"
)

type Wallet struct {
	Id                 uuid.UUID `json:"id"`
	Name               string    `json:"name" validate:"required"`
	UserId             uuid.UUID `json:"userId" validate:"required"`
	value              int64
	CreatedDate        time.Time `json:"createdDate"`
	LastedModifiedDate time.Time `json:"lastedModifiedDate"`
}

var Repository IRepository

func init() {
	Repository = RepositoryImpl{}
}

// NewWallet Constructor
func NewWallet(name string, userId uuid.UUID) *Wallet {
	if userId == uuid.MustParse("00000000-0000-0000-0000-000000000000") {
		return &Wallet{}
	}
	if name == "" {
		return &Wallet{}
	}

	w := new(Wallet)
	w.Name = name
	w.UserId = userId
	w.CreatedDate = time.Now()
	w.LastedModifiedDate = time.Now()
	w.value = 0

	return w
}

func (w Wallet) Save() Wallet {
	newWallet := NewWallet(w.Name, w.UserId)
	Repository.Save(*newWallet)

	return *newWallet
}

func (w Wallet) Update() Wallet {
	currentWallet := Repository.FindById(w.Id)

	if currentWallet.LastedModifiedDate.Before(w.LastedModifiedDate) {
		wallet := w.newWalletWithId()

		return Repository.Update(*wallet)
	}

	return currentWallet
}

func (w Wallet) FindAll() []Wallet {
	return Repository.FindAll()
}

func (w Wallet) FindById(id uuid.UUID) Wallet {
	return Repository.FindById(id)
}

func (w Wallet) GetValue() int64 {
	return w.value
}

func (w Wallet) newWalletWithId() *Wallet {
	wallet := new(Wallet)
	wallet.Id = w.Id
	wallet.UserId = w.UserId
	wallet.Name = w.Name
	wallet.value = w.value
	wallet.CreatedDate = w.CreatedDate
	wallet.LastedModifiedDate = w.LastedModifiedDate
	return wallet
}
