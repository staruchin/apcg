package controllers

import (
	"fmt"
	"log"
	"net/http"

	"goapi/models"
)

func StartServer() {

	http.HandleFunc("/iapm/api/v1/iapm", createIapm)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatalln(err)
	}
}

func createIapm(w http.ResponseWriter, r *http.Request) {
	cmd := "select name from words where id = ?"
	var word string
	err := models.Db.QueryRow(cmd, 1).Scan(&word)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Fprintf(w, word)
}
