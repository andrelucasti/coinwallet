package test

import (
	"coinwallet/wallet"
	"testing"

	"github.com/google/uuid"
)

func TestWalletWhenUserIdIsEmpty(t *testing.T) {
	w := wallet.Wallet{
		Id:   uuid.MustParse("d4f89e35-eae0-4122-bbfe-6dbb7f3bc654"),
		Name: "Hold 2022",
	}

	expected := wallet.WalletApiError{
		HttpErrorCode: 400,
		Message:       "Bad request: userId is empty, please, put any userId",
	}

	_, actual := w.Save()

	if actual != expected {
		t.Error("Expected: ", expected, "Got: ", actual)
	}
}
