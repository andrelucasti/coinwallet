package dbcmanager

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

//TODO Create a file configuration... yml or .env ...
const datasourceName = "user=coinwallet dbname=coinwallet password=coinwallet host=localhost sslmode=disable"
const driverName = "postgres"

type Dbc struct {
	datasourceName string
	driverName     string
}

// NewDbc Constructor
func NewDbc(driverName string, datasourceName string) *Dbc {
	dbc := new(Dbc)
	dbc.driverName = driverName
	dbc.datasourceName = datasourceName

	return dbc
}

var newDbc Dbc

func init() {
	newDbc = *NewDbc(driverName, datasourceName)
}

func OpenConnection() *sql.DB {
	open, err := sql.Open(newDbc.driverName, newDbc.datasourceName)

	if err != nil {
		log.Fatal(" Error to open connection", err.Error())
	}

	return open
}

func CloseConnection(db *sql.DB) {
	err := db.Close()

	if err != nil {
		log.Fatal("Error to close connection")
	}
}
