package models

import "log"

type Word string

func GetAllWords() (words []string, err error) {
	rows, err := Db.Query("select name from words")
	if err != nil {
		log.Fatalln(err)
	}
	for rows.Next() {
		var word string
		err = rows.Scan(&word)
		if err != nil {
			log.Fatalln(err)
		}
		words = append(words, word)
	}
	rows.Close()
	return words, err
}
