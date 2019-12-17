package services

import (
	"github.com/birukbelay/Aprojects/eventListing/entity"
	"github.com/birukbelay/Aprojects/eventListing/events"
)

type EventServiceImpl struct {
	EventRepo events.EventRepository
}

func NewEventSerImp(evRepo events.EventRepository) *EventServiceImpl {
	return &EventServiceImpl{EventRepo: evRepo}
}

func (es *EventServiceImpl) Events() ([]entity.Event, error) {

	events, err := es.EventRepo.Events()

	if err != nil {
		return nil, err
	}

	return events, nil
}
