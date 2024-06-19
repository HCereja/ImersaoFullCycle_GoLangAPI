package domain

import (
	"errors"
	"time"
)

var (
	ErrEventNameRequired    = errors.New("nome do evento é obrigatório")
	ErrInvalidEventDate     = errors.New("data do evento deve ser depois de hoje")
	ErrInvalidEventCapacity = errors.New("capacidade do evento deve ser maior que 0")
	ErrInvalidEventPrice    = errors.New("valor do evento deve ser maior que 0")
)

type Rating string

const (
	RatingLivre Rating = "L"
	Rating10    Rating = "L10"
	Rating12    Rating = "L12"
	Rating14    Rating = "L14"
	Rating16    Rating = "L16"
	Rating18    Rating = "L18"
)

type Event struct {
	ID           string
	Name         string
	Location     string
	Organization string
	Rating       Rating
	Date         time.Time
	ImageURL     string
	Capacity     int
	Price        float64
	PartnerID    int
	Spots        []Spot
	Tickets      []Ticket
}

func (e Event) Validate() error {
	if e.Name == "" {
		return ErrEventNameRequired
	}

	if e.Date.Before(time.Now()) {
		return ErrInvalidEventDate
	}

	if e.Capacity <= 0 {
		return ErrInvalidEventCapacity
	}

	if e.Price <= 0 {
		return ErrInvalidEventPrice
	}

	return nil
}

func (e *Event) AddSpot(name string) (*Spot, error) {
	spot, err := NewSpot(e, name)

	if err != nil {
		return nil, err
	}

	e.Spots = append(e.Spots, *spot)

	return spot, nil
}
