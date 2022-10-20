package controllers

import (
	"fmt"
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

	http.HandleFunc("/iapm/api/v1/chara", chara)

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

func createRand(num int) int {
	return rand.Intn(num)
}

func chara(w http.ResponseWriter, r *http.Request) {
	word := words[createRand(numOfWords)]
	suffix := suffixes[createRand(numOfSuffixes)]
	if !suffix.IsPrefix {
		fmt.Fprintf(w, word+suffix.Name)
	} else {
		fmt.Fprintf(w, suffix.Name+word)
	}
}
