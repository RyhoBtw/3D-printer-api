package db

import (
	"database/sql"
	"fmt"

	"github.com/RyhoBtw/3D-printer-api/log"
	_ "github.com/go-sql-driver/mysql"
)

const (
	DB_HOST = "IP"
	DB_NAME = "3D_Print"
	DB_USER = "root"
	DB_PASS = "Password"
)

type User struct {
	Id        int    `json:"id"`
	Firstname string `form:"firstname"`
	Lastname  string `form:"lastname"`
	Username  string `form:"username"`
	Password  string `form:"password"`
	Email     string `form:"email"`
	TNR       string `form:"tnr"`
	LoginKey  string `json:"loginKey"`
}

func dsn(dbName string) string {
	return fmt.Sprintf("%s:%s@tcp(%s)/%s", DB_USER, DB_PASS, DB_HOST, DB_NAME)
}

func ConnectToDatabase() *sql.DB {
	log.Log().Infoln(DB_HOST)
	db := OpenDB()

	pingErr := db.Ping()
	if pingErr != nil {
		log.Log().Panic(pingErr)
	}

	defer db.Close()

	return db
}

func OpenDB() *sql.DB {
	db, err := sql.Open("mysql", dsn(""))
	if err != nil {
		log.Log().Infof("Error %s when opening DB\n", err)
	}
	return db
}

func GetInfo() {
	db := OpenDB()
	defer db.Close()

	//row := db.QueryRow("SELECT * FROM user")
	//log.Log().Infoln(row)
	rows, err := db.Query("SELECT * FROM user")
	if err != nil {
		panic(err.Error())
	}

	for rows.Next() {
		var user User
		_ = rows.Scan(&user.Id, &user.Username, &user.Password, &user.LoginKey)
		/*if err != nil {
			panic(err.Error())
		}*/

		fmt.Printf("%v\n", user)
		//json.NewEncoder(w).Encode(user)
	}

}
