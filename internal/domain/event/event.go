package event

type Event struct {
	ID               uint    `db:"id,omitempty"`
	Title            string  `db:"Title"`
	ShortDescription string  `db:"Short Description"`
	Description      string  `db:"Description"`
	Longitude        float64 `db:"Longitude"`
	Latitude         float64 `db:"Latitude"`
	Images           string  `db:"Images"`
	Preview          string  `db:"Preview"`
	Date             string  `db:"Date"`
}
