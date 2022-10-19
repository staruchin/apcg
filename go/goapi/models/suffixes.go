package models

import "time"

type Suffix struct {
	ID        int
	Name      string
	IsPrefix  bool
	CreatedAt time.Time
}
