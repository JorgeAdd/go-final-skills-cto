package database

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

func GetDB() *sql.DB {
	psqlInfo := fmt.Sprintf("host=bitsodb-instance-1.cfchj75aj0uh.us-east-1.rds.amazonaws.com port=5432 user=postgres " +
		"password=password dbname=BitsoDB sslmode=disable")
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	return db
}
