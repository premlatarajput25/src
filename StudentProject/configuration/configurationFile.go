package configuration

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func GetDBConnection() (db *sql.DB, err error) {

	dbDriver := "mysql"
	ServerName := "localhost"
	dbPort := "3306"
	dbUser := "premlata"
	dbPassword := "premlata"
	dbName := "studentdb"
	// db, dberr := sql.Open("mysql",   "root:root@tcp(127.0.0.1:3306)/studentdb)
	db, err = sql.Open(dbDriver, dbUser+":"+dbPassword+"@tcp("+ServerName+":"+dbPort+")/"+dbName)
	// fmt.Println("SEE " + err.Error())
	return
}
