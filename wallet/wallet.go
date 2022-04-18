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
	repository         Repository
}

//Constructor
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
	w.repository.Save(*newWallet)

	return *newWallet
}

func (w Wallet) FindAll() []Wallet {
	return w.repository.FindAll()
}

func (w Wallet) GetValue() int64 {
	return w.value
}
