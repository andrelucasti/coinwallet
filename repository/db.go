package repository

import "database/sql"

//TODO Testar
const datasourceName = "user=coinwallet dbname=coinwallet password=coinwallet host=localhost sslmode=disable"
const driverName = "postgres"

type dbc struct {
	datasourceName string
	driverName     string
}

func init() {
	fetchDataDataSource()
}

func (d dbc) OpenConnection() {
	sql.Open(d.driverName, d.datasourceName)
}

func fetchDataDataSource() {
	d := dbc{}
	d.driverName = driverName
	d.datasourceName = datasourceName
}
