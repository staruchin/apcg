package models

import (
	"log"
)

type Suffix struct {
	Name     string
	IsPrefix bool
	Weight   int
}

func GetAllSuffixes() (suffixes []Suffix, err error) {
	rows, err := Db.Query("select name, is_prefix, weight from suffixes")
	if err != nil {
		log.Fatalln(err)
	}
	for rows.Next() {
		var suffix Suffix
		err = rows.Scan(&suffix.Name, &suffix.IsPrefix, &suffix.Weight)
		if err != nil {
			log.Fatalln(err)
		}
		suffixes = append(suffixes, suffix)
	}
	rows.Close()
	return suffixes, err
}
