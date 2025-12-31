package db

import (
	"fmt"

	"blog/config"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func GetConnectionString(cnf *config.DBConfig) string {
	connString := fmt.Sprintf("user=%s password=%s host=%s port=%d dbname=%s",
		cnf.DBUser,
		cnf.DBPassword,
		cnf.DBHost,
		cnf.DBPort,
		cnf.DBName,
	)

	if !cnf.DBSSLMode {
		connString += " sslmode=disable"
	}

	return connString
	// return "user=postgres password=password host=localhost port=5432 dbname=blog sslmode=disable"
}

func NewConnection(cnf *config.DBConfig) (*sqlx.DB, error) {
	dbSource := GetConnectionString(cnf)
	dbCon, err := sqlx.Connect("postgres", dbSource)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return dbCon, nil
}
