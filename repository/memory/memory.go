package memory

import (
	"github.com/bemsk/parky/parking"
	"github.com/bemsk/parky/repository"
)

type Memory struct {
	parking *parking.Parking
	tickets map[string]*parking.Ticket
}

func New() *Memory {
	p := parking.New()
	ts := make(map[string]*parking.Ticket)

	return &Memory{p, ts}
}

func (m *Memory) Save(p *parking.Parking) {
	m.parking = p
}

func (m *Memory) Load() *parking.Parking {
	return m.parking
}

func (m *Memory) GetTicket(registrationNumber string) (*parking.Ticket, error) {
	ticket, prs := m.tickets[registrationNumber]

	if prs {
		return ticket, nil
	}

	return nil, repository.ErrRepositoryTicketNotRegistered
}

func (m *Memory) GetTickets(color parking.ReadableColor) []*parking.Ticket {
	tickets := []*parking.Ticket{}

	for _, t := range m.tickets {
		if t.Color.ReadableName() == color.ReadableName() {
			tickets = append(tickets, t)
		}
	}

	return tickets
}

func (m *Memory) AddTicket(p *parking.Parking, t *parking.Ticket) error {
	if _, prs := m.tickets[t.RegistrationNumber]; prs {
		return repository.ErrRepositoryTicketAlreadyRegistered
	}
	m.parking = p
	m.tickets[t.RegistrationNumber] = t

	return nil
}

func (m *Memory) RemoveTicket(p *parking.Parking, t *parking.Ticket) error {
	if _, prs := m.tickets[t.RegistrationNumber]; !prs {
		return repository.ErrRepositoryTicketNotRegistered
	}

	m.parking = p
	delete(m.tickets, t.RegistrationNumber)

	return nil
}
