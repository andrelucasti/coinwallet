package wallet

import (
	"coinwallet/dbcmanager"
	"database/sql"
	"github.com/google/uuid"
	"log"
)

type IRepository interface {
	Save(w Wallet)
	Update(w Wallet) Wallet
	FindAll() []Wallet
	FindById(id uuid.UUID) Wallet
}

type RepositoryImpl struct{}

func (r RepositoryImpl) Save(w Wallet) {
	db := dbcmanager.OpenConnection()
	query := "INSERT INTO WALLET (name, user_id, value, created_date, last_modified_date) VALUES($1, $2, $3, $4, $5)"

	if _, err := db.Query(query, w.Name, w.UserId, w.GetValue(), w.CreatedDate, w.LastedModifiedDate); err != nil {
		log.Fatal("Error to persist wallet", err)
	}

	defer dbcmanager.CloseConnection(db)
}

func (r RepositoryImpl) FindAll() []Wallet {
	db := dbcmanager.OpenConnection()
	defer dbcmanager.CloseConnection(db)
	return fetchWallets(db)
}

func (r RepositoryImpl) DeleteAll() {
	db := dbcmanager.OpenConnection()
	_, err := db.Query("DELETE FROM wallet")

	if err != nil {
		log.Fatal("Error to delete all wallets", err)
	}

	defer dbcmanager.CloseConnection(db)
}

func (r RepositoryImpl) FindById(id uuid.UUID) Wallet {
	return Wallet{Name: "mi ovito"}
}

func (r RepositoryImpl) Update(wallet Wallet) Wallet {
	return wallet
}

func fetchWallets(db *sql.DB) []Wallet {
	query := "SELECT * FROM WALLET"
	rows, err := db.Query(query)
	if err != nil {
		log.Fatal(" Error to get all Wallets", err.Error())
	}

	return buildWallets(rows)
}

func buildWallets(rows *sql.Rows) []Wallet {
	var wallets []Wallet
	var wallet Wallet

	for rows.Next() {
		if err := rows.Scan(
			&wallet.Id,
			&wallet.Name,
			&wallet.UserId,
			&wallet.value,
			&wallet.CreatedDate,
			&wallet.LastedModifiedDate); err != nil {

			log.Fatal(" Error to scanner the wallets", err.Error())
		}

		wallets = append(wallets, wallet)

	}
	return wallets
}
