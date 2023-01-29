package parking

import (
	"errors"
)

var (
	ErrParkingSlotUnavailable = errors.New("[parking] slot unavailable")
	ErrParkingInputRequired = errors.New("[parking] value cannot be empty")
	ErrParkingCapacityMustBeGreaterThanOne = errors.New("[parking] capacity must be greater than one")
	ErrParkingInvalidTicket = errors.New("[parking] invalid ticket")
)

const (
	occupied slot = "occupied"
	empty slot = "empty"
)

type ReadableColor interface {
	ReadableName() string
}

type slot string

type Car struct {
	RegistrationNumber string
	Color ReadableColor
}

type Ticket struct {
	*Car
	SlotNumber int
}

type Parking struct {
	capacity int
	slots []slot
}

func New() *Parking {
	slots := make([]slot, 0)

	return &Parking{0, slots}
}

func (p *Parking) SetCapacity(capacity int) error {
	if (capacity < 1) {
		return ErrParkingCapacityMustBeGreaterThanOne
	}

	p.capacity = capacity
	slots := make([]slot, capacity)

	for i := range slots {
		slots[i] = empty
	}

	p.slots = slots

	return nil
}

func (p *Parking) IssueTicket(registrationNumber string, color ReadableColor) (*Ticket, error) {
	if registrationNumber == "" || color.ReadableName() == "" {
		return nil, ErrParkingInputRequired
	}

	car := &Car{registrationNumber, color}
	slotNumber, err := p.nearestEmptySlot()

	if err != nil {
		return nil, err
	}

	p.occupySlot(slotNumber)

	return &Ticket{ car, slotNumber }, nil
}

func (p *Parking) Vacate(t *Ticket) error {
	if t.SlotNumber >= len(p.slots) || p.slots[t.SlotNumber] == empty {
		return ErrParkingInvalidTicket
	}

	p.slots[t.SlotNumber] = empty

	return nil
}

func (p *Parking) nearestEmptySlot() (int, error) {
	for i, s := range p.slots {
		if s == empty {
			return i, nil
		}
	}

	return -1, ErrParkingSlotUnavailable
}

func (p *Parking) occupySlot(slotNumber int) {
	p.slots[slotNumber] = occupied
}