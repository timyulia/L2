package repository

import (
	"calendar"
	"time"
)

type Cal interface {
	CreateEvent(event calendar.Event) error
	UpdateEvent(event calendar.Event) error
	DeleteEvent(id int) error
	ForDay(id int, date time.Time) ([]calendar.Event, error)
	ForWeek(id int, date time.Time) ([]calendar.Event, error)
	ForMonth(id int, date time.Time) ([]calendar.Event, error)
}

type Repository struct {
	Cal
}

func NewRepository() *Repository {
	return &Repository{Cal: NewCalMem()}
}
