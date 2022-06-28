package domain

import "time"

type Events struct {
	Id               uint      `db:"id,omitempty"`
	Title            string    `db:"Title"`
	ShortDescription string    `db:"ShortDescription"`
	Description      string    `db:"Description"`
	Longitude        float64   `db:"Longitude"`
	Latitude         float64   `db:"Latitude"`
	Images           string    `db:"Images"`
	Preview          string    `db:"Preview"`
	Date             string    `db:"Date"`
	IsEnded          bool      `db:"isEnded"`
	DeletedDate      time.Time `db:"deletedDate"`
}
