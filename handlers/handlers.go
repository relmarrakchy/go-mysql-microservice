package handlers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"gosql/models"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func init() {
	var err error
	db, err = sql.Open("mysql", "root:@tcp(localhost:3306)/training")
	if err != nil {
		fmt.Println("Error while connecting to database")
	}

	err = db.Ping()
	if err != nil {
		fmt.Println("Error while pinging ...")
	}
}

func GetUsers(w http.ResponseWriter, r *http.Request) {
	result, err := db.Query("SELECT * FROM users")
	if err != nil {
		fmt.Println("Error while fetching data ...")
	}

	var users []models.User

	for result.Next() {
		var user models.User
		err = result.Scan(&user.IDuser, &user.Username, &user.Email)
		if err != nil {
			panic(err.Error())
		}
		users = append(users, user)
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}

func InsertUser(username string, email string) {
	query := "INSERT INTO users VALUES (null, ?, ?)"
	_, err := db.Exec(query, username, email)
	if err != nil {
		panic(err)
	}
}
