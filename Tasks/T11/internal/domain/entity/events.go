package entity

import "time"

type Event struct {
	Datetime    time.Time
	Title       string
	Description string
}
