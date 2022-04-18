package wallet

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

type Repository struct{}

//TODO Create a file configuration... yml or .env ...
const datasourceName = "user=coinwallet dbname=coinwallet password=coinwallet host=localhost sslmode=disable"
const driverName = "postgres"

func (r Repository) Save(w Wallet) {
	// TODO refactor
	if db, err := sql.Open(driverName, datasourceName); err == nil {
		query := "INSERT INTO WALLET (name, user_id, value, created_date, last_modified_date) VALUES($1, $2, $3, $4, $5)"

		if _, err := db.Query(query, w.Name, w.UserId, w.GetValue(), w.CreatedDate, w.LastedModifiedDate); err != nil {
			log.Fatal("Error to persist wallet", err)
		}

		defer db.Close()

	} else {
		log.Fatal(" Error to open connection", err.Error())

	}

}

func (r Repository) FindAll() []Wallet {
	if db, err := sql.Open(driverName, datasourceName); err == nil {
		log.Println("Opened connection")

		return fetchWallets(db)

	} else {
		log.Fatal(" Error to open connection with database", err.Error())
		return nil
	}
}

func (r Repository) DeleteAll() {
	if db, err := sql.Open(driverName, datasourceName); err == nil {
		_, err := db.Query("DELETE FROM wallet")

		if err != nil {
			log.Fatal("Error to delete all wallets", err)
		}

		defer db.Close()

	} else {
		log.Fatal(" Error to open connection with database", err.Error())
	}

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
