package app

import (
	"errors"

	"github.com/bemsk/parky/parking"
)

var (
	ErrAppCarAlreadyRegistered = errors.New("[app] car already registered")
)

type Repository interface {
	Save(p *parking.Parking)
	Load() *parking.Parking
	GetTicket(registrationNumber string) (*parking.Ticket, error)
	GetTickets(color parking.ReadableColor) []*parking.Ticket
	AddTicket(p *parking.Parking, t *parking.Ticket) error
	RemoveTicket(p *parking.Parking, t *parking.Ticket) error
}

type App struct {
	repository Repository
	colors []parking.ReadableColor
}

func New(r Repository, c []parking.ReadableColor) *App {
	return &App{r, c}
}

func (a *App) SetParkingCapacity(capacity int) {
	p := a.repository.Load()

	p.SetCapacity(capacity)

	a.repository.Save(p)
} 

func (a *App) OrderTicket(registrationNumber string, color parking.ReadableColor) (*parking.Ticket, error) {
	p := a.repository.Load()

	ticket, _ := a.repository.GetTicket(registrationNumber)

	if ticket != nil {
		return nil, ErrAppCarAlreadyRegistered
	}

	t, err := p.IssueTicket(registrationNumber, color)

	if err != nil {
		return nil, err
	}
	
	err = a.repository.AddTicket(p, t)

	if err != nil {
		return nil, err
	}

	return t, nil
}

func (a *App) RetractTicket(registrationNumber string) error {
	p := a.repository.Load()
	t, err := a.repository.GetTicket(registrationNumber)

	if err != nil {
		return err
	}

	err = p.Vacate(t)

	if err != nil {
		return err
	}

	err = a.repository.RemoveTicket(p, t)

	if err != nil {
		return err
	}

	return nil
}

func (a *App) SlotNumber(registrationNumber string) (int, error) {
	t, err := a.repository.GetTicket(registrationNumber)

	if err != nil {
		return -1, err
	}

	return t.SlotNumber, nil
}

func (a *App) RegistrationNumbersByColor(color parking.ReadableColor) []string {
	var rns []string

	tickets := a.repository.GetTickets(color)

	for _, t := range tickets {
		rns = append(rns, t.RegistrationNumber)
	}

	return rns
}

func (a *App) SlotNumbersByColor(color parking.ReadableColor) []int {
	var sns []int

	tickets := a.repository.GetTickets(color)

	for _, t := range tickets {
		sns = append(sns, t.SlotNumber)
	}

	return sns
}

func (a *App) ListAvailableColors() []parking.ReadableColor {
	return a.colors
}