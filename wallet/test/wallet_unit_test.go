package test

import (
	"coinwallet/wallet"
	"coinwallet/wallet/test/mocks"
	"testing"
	"time"

	"github.com/google/uuid"
)

func TestWalletWhenUserIdIsEmpty(t *testing.T) {
	wallet.Repository = mocks.RepositoryMock{}
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
	wallet.Repository = mocks.RepositoryMock{}
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
	wallet.Repository = mocks.RepositoryMock{}
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

func TestWalletWhenCreatedTheValueIsEqualZero(t *testing.T) {
	wallet.Repository = mocks.RepositoryMock{}
	w := wallet.Wallet{
		Name:               "My Wallet From Crypto",
		UserId:             uuid.MustParse("258BAE13-F477-4F96-9C7C-D9124A10A53E"),
		CreatedDate:        time.Now(),
		LastedModifiedDate: time.Now(),
	}

	actual := w.Save()

	if actual.GetValue() != 0 {
		t.Error("Expected: ", 0, "Got: ", actual)
	}
}

func TestNotUpdateWalletNameWhenCurrentLastModifiedDateIsBeforeThanLastModifiedDate(t *testing.T) {
	wallet.Repository = mocks.RepositoryMock{}

	currentWalletName := "Games Token"
	mocks.FindByIdMock = func(id uuid.UUID) wallet.Wallet {
		return wallet.Wallet{
			Id:                 uuid.MustParse("1B425A4E-8BDC-419D-9B9B-7E2C090A2E49"),
			Name:               currentWalletName,
			UserId:             uuid.MustParse("258BAE13-F477-4F96-9C7C-D9124A10A53E"),
			CreatedDate:        time.Now().AddDate(2022, 05, 15),
			LastedModifiedDate: time.Now().Add(time.Hour + (time.Minute * 10)),
		}
	}

	newWalletName := "My Wallet From Crypto"
	w := wallet.Wallet{
		Id:                 uuid.MustParse("1B425A4E-8BDC-419D-9B9B-7E2C090A2E49"),
		Name:               newWalletName,
		UserId:             uuid.MustParse("258BAE13-F477-4F96-9C7C-D9124A10A53E"),
		CreatedDate:        time.Now().AddDate(2022, 05, 15),
		LastedModifiedDate: time.Now().Add(time.Hour * -1),
	}

	actual := w.Update()

	if actual.Name != currentWalletName {
		t.Error("Expected: ", currentWalletName, "Got: ", actual.Name)
	}
}

func TestUpdateWalletNameWhenCurrentLastModifiedDateIsBeforeThanLastModifiedDate(t *testing.T) {
	wallet.Repository = mocks.RepositoryMock{}

	currentWalletName := "Games Token"
	mocks.FindByIdMock = func(id uuid.UUID) wallet.Wallet {
		return wallet.Wallet{
			Id:                 uuid.MustParse("1B425A4E-8BDC-419D-9B9B-7E2C090A2E49"),
			Name:               currentWalletName,
			UserId:             uuid.MustParse("258BAE13-F477-4F96-9C7C-D9124A10A53E"),
			CreatedDate:        time.Now().AddDate(2022, 05, 15),
			LastedModifiedDate: time.Now().Add(time.Hour * -1),
		}
	}

	newWalletName := "My Wallet From Crypto"
	w := wallet.Wallet{
		Id:                 uuid.MustParse("1B425A4E-8BDC-419D-9B9B-7E2C090A2E49"),
		Name:               newWalletName,
		UserId:             uuid.MustParse("258BAE13-F477-4F96-9C7C-D9124A10A53E"),
		CreatedDate:        time.Now().AddDate(2022, 05, 15),
		LastedModifiedDate: time.Now().Add(time.Hour + (time.Minute * 10)),
	}

	actual := w.Update()

	if actual.Name != newWalletName {
		t.Error("Expected: ", newWalletName, "Got: ", actual.Name)
	}
}

func TestUpdateWalletWhenCurrentLastModifiedDateIsAfterThanLastModifiedDate(t *testing.T) {
	wallet.Repository = mocks.RepositoryMock{}
	mocks.FindByIdMock = func(id uuid.UUID) wallet.Wallet {
		return wallet.Wallet{
			Id:                 uuid.MustParse("1B425A4E-8BDC-419D-9B9B-7E2C090A2E49"),
			Name:               "Games Token",
			UserId:             uuid.MustParse("258BAE13-F477-4F96-9C7C-D9124A10A53E"),
			CreatedDate:        time.Now().AddDate(2022, 05, 15),
			LastedModifiedDate: time.Now().Add(time.Hour * -1),
		}
	}

	w := wallet.Wallet{
		Id:                 uuid.MustParse("1B425A4E-8BDC-419D-9B9B-7E2C090A2E49"),
		Name:               "My Wallet From Crypto",
		UserId:             uuid.MustParse("258BAE13-F477-4F96-9C7C-D9124A10A53E"),
		CreatedDate:        time.Now().AddDate(2022, 05, 15),
		LastedModifiedDate: time.Now().Add(time.Hour + (time.Minute * 10)),
	}
	actual := w.Update()

	if actual != w {
		t.Error("Expected: ", w, "Got: ", actual)
	}
}
