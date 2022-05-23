package event

import (
	"github.com/upper/db/v4"
	"github.com/upper/db/v4/adapter/postgresql"
	"log"
)

var settings = postgresql.ConnectionURL{
	Database: `training`,
	Host:     `localhost:54322`,
	User:     `postgres`,
	Password: `root`,
}

type Repository interface {
	FindAll() ([]Event, error)
	FindOne(id uint64) (*Event, error)
}

const EventsCount uint64 = 8

type repository struct {
	//ID               uint    `db:"id,omitempty"`
	//Title            string  `db:"Title"`
	//ShortDescription string  `db:"Short Description"`
	//Description      string  `db:"Description"`
	//Longitude        float64 `db:"Longitude"`
	//Latitude         float64 `db:"Latitude"`
	//Images           string  `db:"Images"`
	//Preview          string  `db:"Preview"`
	//Date             string  `db:"Date"`
}

func NewRepository() Repository {
	return &repository{}
}

func (r *repository) FindAll() ([]Event, error) {
	events := make([]Event, EventsCount)

	sess, err := postgresql.Open(settings)
	if err != nil {
		log.Fatal("Open: ", err)
	}
	defer sess.Close()
	eventCol := sess.Collection("events")
	err = eventCol.Find().All(&events)
	return events, nil
}

func (r *repository) FindOne(id uint64) (*Event, error) {
	var entity Event
	sess, err := postgresql.Open(settings)
	if err != nil {
		log.Fatal("Open: ", err)
	}
	defer sess.Close()
	eventCol := sess.Collection("events")
	err = eventCol.Find(
		db.Cond{"id": id},
	).One(&entity)
	return &entity, nil
}
