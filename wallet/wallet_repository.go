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
	query := "INSERT INTO WALLET (name, user_id, value, created_date, last_modified_date) VALUES($1, $2, $3, $4, $5)"

	db := dbcmanager.OpenConnection()
	defer dbcmanager.CloseConnection(db)

	if _, err := db.Query(query, w.Name, w.UserId, w.GetValue(), w.CreatedDate, w.LastedModifiedDate); err != nil {
		log.Fatal("Error to persist wallet", err)
	}

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
		log.Fatal("Error to delete all wallets", err.Error())
	}

	defer dbcmanager.CloseConnection(db)
}

func (r RepositoryImpl) FindById(id uuid.UUID) Wallet {
	query := "SELECT  * FROM wallet WHERE uuid = $1"

	db := dbcmanager.OpenConnection()
	defer dbcmanager.CloseConnection(db)

	var wallet Wallet
	if rows, err := db.Query(query, id); err == nil {
		if rows.Next() {
			if err := rows.Scan(
				&wallet.Id,
				&wallet.Name,
				&wallet.UserId,
				&wallet.value,
				&wallet.CreatedDate,
				&wallet.LastedModifiedDate); err != nil {
				log.Fatalln(" Error to scanner the wallet. WalletId: ", id, "\n error: ", err.Error())
			}
		}
	} else {
		log.Fatalln(" Error to get wallet by WalletId", id, "\n error: ", err.Error())
	}

	return wallet
}

func (r RepositoryImpl) Update(wallet Wallet) Wallet {
	query := "UPDATE wallet SET name = $2, last_modified_date = now() WHERE uuid = $1"

	db := dbcmanager.OpenConnection()
	defer dbcmanager.CloseConnection(db)

	if _, err := db.Query(query, wallet.Id, wallet.Name); err != nil {
		log.Fatalln("Error to update the wallet. Wallet ID:", wallet.Id, "Error: ", err.Error())
	}

	return r.FindById(wallet.Id)
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
