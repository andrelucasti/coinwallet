package wallet

import (
	"time"

	"github.com/google/uuid"
)

type Wallet struct {
	Id                 uuid.UUID `json:"id"`
	Name               string    `json:"name"`
	UserId             uuid.UUID `json:"userId"`
	CreatedDate        time.Time `json:"createdDate"`
	LastedModifiedDate time.Time `json:"lastedModifiedDate"`
}

type WalletApiError struct {
	HttpErrorCode int
	Message       string
}

func (w Wallet) Save() (Wallet, WalletApiError) {
	var wApiError = WalletApiError{}

	if w.UserId == uuid.MustParse("00000000-0000-0000-0000-000000000000") {
		wApiError.HttpErrorCode = 400
		wApiError.Message = "Bad request: userId is empty, please, put any userId"
	}

	return (Wallet{}), (wApiError)
}
