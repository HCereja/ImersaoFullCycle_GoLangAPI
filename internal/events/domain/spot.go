package domain

import (
	"errors"

	"github.com/google/uuid"
)

var (
	ErrSpotNameRequired    = errors.New("nome do lugar é obrigatório")
	ErrSpotNameLength      = errors.New("nome do lugar deve ter pelo menos 2 caracteres")
	ErrSpotNameStart       = errors.New("nome do lugar deve começar com uma letra")
	ErrSpotNameEnd         = errors.New("nome do lugar deve terminar com um número")
	ErrInvalidSpotNumber   = errors.New("número do lugar inválido")
	ErrSpotNotFound        = errors.New("lugar não foi encontrado")
	ErrSpotAlreadyReserved = errors.New("lugar já foi reservado")
)

type SpotStatus string

const (
	SpotStatusAvailable SpotStatus = "available"
	SpotStatusSold      SpotStatus = "sold"
)

type Spot struct {
	ID       string
	EventID  string
	Name     string
	Status   SpotStatus
	TicketID string
}

func (s Spot) Validate() error {
	if s.Name == "" {
		return ErrSpotNameRequired
	}

	if len(s.Name) < 2 {
		return ErrSpotNameLength
	}

	if s.Name[0] < 'A' || s.Name[0] > 'Z' {
		return ErrSpotNameStart
	}

	if s.Name[len(s.Name)-1] < '0' || s.Name[len(s.Name)-1] > '9' {
		return ErrSpotNameEnd
	}

	return nil
}

func NewSpot(event *Event, name string) (*Spot, error) {
	spot := &Spot{
		ID:      uuid.New().String(),
		EventID: event.ID,
		Name:    name,
		Status:  SpotStatusAvailable,
	}

	if err := spot.Validate(); err != nil {
		return nil, err
	}
	return spot, nil
}

func (s *Spot) Reserve(ticketId string) error {
	if s.Status == SpotStatusSold {
		return ErrSpotAlreadyReserved
	}

	s.Status = SpotStatusSold
	s.TicketID = ticketId
	return nil
}
