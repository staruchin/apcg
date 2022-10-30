package controllers

import (
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
	"time"

	"goapi/models"
)

var suffixes []models.Suffix
var numOfSuffixes int
var words []string
var numOfWords int

func StartServer() {

	initialize()

	http.HandleFunc("/apcg/api/v1/character", character)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatalln(err)
	}
}

func initialize() {

	ss, err := models.GetAllSuffixes()
	if err != nil {
		log.Fatalln(err)
	}
	suffixes = []models.Suffix{}
	for _, s := range ss {
		for i := 0; i < s.Weight; i++ {
			suffixes = append(suffixes, s)
		}
	}
	numOfSuffixes = len(suffixes)

	words, err = models.GetAllWords()
	if err != nil {
		log.Fatalln(err)
	}
	numOfWords = len(words)

	rand.Seed(time.Now().UnixNano())
}

func character(w http.ResponseWriter, r *http.Request) {

	word := words[rand.Intn(numOfWords)]
	suffix := suffixes[rand.Intn(numOfSuffixes)]

	w.Header().Set("Access-Control-Allow-Headers", "*")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)

	var name string
	if !suffix.IsPrefix {
		name = word + suffix.Name
	} else {
		name = suffix.Name + word
	}

	if err := json.NewEncoder(w).Encode(map[string]string{"name": name}); err != nil {
		log.Println(err)
	}
}
