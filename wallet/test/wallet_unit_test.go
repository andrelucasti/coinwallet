package test

import (
	"coinwallet/wallet"
	"testing"
	"time"

	"github.com/google/uuid"
)

func TestWalletWhenUserIdIsEmpty(t *testing.T) {
	w := wallet.Wallet{
		Name: "Hold 2022",
	}

	expected := wallet.Wallet{}
	actual := w.Save()

	if actual != expected {
		t.Error("Expected: ", expected, "Got: ", actual)
	}
}

func TestWalletWhenNameIsEmpty(t *testing.T) {
	w := wallet.Wallet{
		UserId: uuid.MustParse("258BAE13-F477-4F96-9C7C-D9124A10A53E"),
	}

	expected := wallet.Wallet{}
	actual := w.Save()

	if actual != expected {
		t.Error("Expected: ", expected, "Got: ", actual)
	}
}

func TestSaveWallet(t *testing.T) {
	w := wallet.Wallet{
		Name:   "CryptoGames",
		UserId: uuid.MustParse("258BAE13-F477-4F96-9C7C-D9124A10A53E"),
	}

	actual := w.Save()

	expected := wallet.Wallet{
		Name:               "CryptoGames",
		UserId:             uuid.MustParse("258BAE13-F477-4F96-9C7C-D9124A10A53E"),
		CreatedDate:        time.Now(),
		LastedModifiedDate: time.Now(),
	}

	if actual.UserId != expected.UserId ||
		actual.Name != expected.Name {
		t.Error("Expected: ", expected, "Got: ", actual)
	}
}
