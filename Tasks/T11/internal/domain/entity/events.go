package entity

import (
	"fmt"
	"time"
)

type Event struct {
	Datetime    time.Time
	Title       string
	Description string
}

func (e Event) String() string {
	return fmt.Sprintf(
		"{\"title\":\"%s\", \"description\":\"%s\", \"datetime\":\"%s\"}",
		e.Title, e.Description, e.Datetime.Format(time.RFC3339))
}
