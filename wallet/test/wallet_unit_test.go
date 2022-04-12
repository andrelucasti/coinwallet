package test

import (
	"coinwallet/wallet"
	"testing"

	"github.com/google/uuid"
)

func TestWalletWhenUserIdIsEmpty(t *testing.T) {
	w := wallet.Wallet{
		Id:   uuid.MustParse("D3CFD8AC-CB66-47AB-AC60-AE6D64B89493"),
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
		Id:     uuid.MustParse("D3CFD8AC-CB66-47AB-AC60-AE6D64B89493"),
		UserId: uuid.MustParse("258BAE13-F477-4F96-9C7C-D9124A10A53E"),
	}

	expected := wallet.Wallet{}
	actual := w.Save()

	if actual != expected {
		t.Error("Expected: ", expected, "Got: ", actual)
	}
}
