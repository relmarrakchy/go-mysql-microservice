package main

import (
	"fmt"
	"gosql/handlers"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/users", handlers.GetUsers).Methods("GET")

	port := 8080
	fmt.Printf("The go server is running on port %d ...\n", port)
	log.Fatal(http.ListenAndServe(":8080", r))
}
