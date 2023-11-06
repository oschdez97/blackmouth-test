package test

import (
	"database/sql"
	"fmt"
	"log"
	_ "github.com/lib/pq"
)

var (
	UNAMEDB string = "postgres"
	PASSDB  string = "postgres123"
	HOSTDB  string = "postgres"
	DBNAME  string = "bmdb"
)

var DB *sql.DB

func init() {
	connStr := fmt.Sprintf("postgres://%s:%s@%s/%s?sslmode=disable", UNAMEDB, PASSDB, HOSTDB, DBNAME)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	DB = db
}