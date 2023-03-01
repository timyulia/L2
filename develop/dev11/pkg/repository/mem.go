package repository

import (
	"calendar"
	"errors"
	"sync"
	"time"
)

type CalMem struct {
	//serial int
	mutex  *sync.RWMutex
	events map[int]calendar.Event
}

func NewCalMem() *CalMem {
	return &CalMem{mutex: &sync.RWMutex{}, events: make(map[int]calendar.Event)}
}

func (m *CalMem) CreateEvent(event calendar.Event) error {
	m.mutex.Lock()
	m.events[event.ID] = event
	m.mutex.Unlock()
	return nil
}

func (m *CalMem) UpdateEvent(event calendar.Event) error {
	if _, ok := m.events[event.ID]; !ok {
		return errors.New("no such event")
	}
	m.events[event.ID] = event
	return nil
}

func (m *CalMem) DeleteEvent(id int) error {
	if _, ok := m.events[id]; !ok {
		return errors.New("no such event")
	}
	delete(m.events, id)
	return nil
}

func (m *CalMem) ForDay(id int, date time.Time) ([]calendar.Event, error) {
	events := make([]calendar.Event, 0)
	m.mutex.RLock()
	for _, val := range m.events {
		if val.Date.Equal(date) && val.UserID == id {
			events = append(events, val)
		}
	}
	m.mutex.RUnlock()
	return events, nil
}

func (m *CalMem) ForWeek(id int, date time.Time) ([]calendar.Event, error) {
	events := make([]calendar.Event, 0)
	m.mutex.RLock()

	for _, event := range m.events {
		difference := event.Date.Sub(date)
		if event.UserID == id && difference >= 0 && difference <= time.Duration(7*24)*time.Hour { //будущие события на неделю
			events = append(events, event)
		}
	}
	m.mutex.RUnlock()
	return events, nil
}

func (m *CalMem) ForMonth(id int, date time.Time) ([]calendar.Event, error) {
	events := make([]calendar.Event, 0)
	m.mutex.RLock()
	for _, event := range m.events {
		if date.Year() == event.Date.Year() && date.Month() == event.Date.Month() && event.UserID == id {
			events = append(events, event)
		}
	}
	m.mutex.RUnlock()
	return events, nil
}
