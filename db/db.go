package db

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func ConectaBancoDados() *sql.DB {
	db, err := sql.Open("mysql", "root@tcp(loja_camisetas_mysql:3306)/loja_camisetas")
	if err != nil {
		panic(err.Error())
	}
	return db
}
