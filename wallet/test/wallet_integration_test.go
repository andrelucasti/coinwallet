package test

import (
	"coinwallet/wallet"
	"testing"
	"time"

	"github.com/google/uuid"
)

const InvalidId = "00000000-0000-0000-0000-000000000000"

func init() {

}

func TestSavePersistWallet(t *testing.T) {
	name := "CryptoGames"
	userId := uuid.MustParse("258BAE13-F477-4F96-9C7C-D9124A10A53E")
	createdDate := time.Now()
	lastModifiedDate := time.Now()

	w := wallet.Wallet{
		Name:               name,
		UserId:             userId,
		CreatedDate:        createdDate,
		LastedModifiedDate: lastModifiedDate,
	}

	w.Save()
	actual := w.FindAll()[0]

	if actual.Id == uuid.MustParse(InvalidId) {
		t.Error("Expected valid ID", "Got: ", actual.Id)
	}

	if name != actual.Name {
		t.Error("Expected: ", name, "Got: ", actual.Name)
	}

	if userId != actual.UserId {
		t.Error("Expected: ", userId, "Got: ", actual.UserId)
	}

}
