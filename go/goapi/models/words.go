package models

import "time"

type Word struct {
	ID        int
	Name      string
	CreatedAt time.Time
}
