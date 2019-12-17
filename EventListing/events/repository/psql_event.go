package repository

import (
	"database/sql"
	"errors"

	"github.com/birukbelay/Aprojects/eventListing/entity"
)

// EventRepoImpl implements the event.EventRepository interface
type EventRepoImpl struct {
	conn *sql.DB
	
}

// NewEventRepoImp will create an object of EventRepoImpl
func NewEventRepoImp(con *sql.DB) *EventRepoImpl {
	return &EventRepoImpl{conn: con}
}

// Events returns all Events from the database
func (eri *EventRepoImpl) Events() ([]entity.Event, error) {

	rows, err := eri.conn.Query("SELECT * FROM events;")
	if err != nil {
		return nil, errors.New("Could not query the database")
	}
	defer rows.Close()

	events := []entity.Event{}

	for rows.Next() {
		event := entity.Event{}
		err := rows.Scan(&event.EId, &event.Name, &event.Details, &event.Image, &event.City, &event.Country, &event.Place, &event.Coordinates, &event.host, &event.IsPassed, &event.Rating, &event.PostedDate, &event.price)

		if err != nil {
			return nil, err
		}
		events = append(events, event)
	}
	return events, nil
}

//Event ..returns single event
func (eri *EventRepoImpl) Event(id int) (entity.Event, error) {

	row := eri.conn.QueryRow("SELECT * FROM events WHERE id = $1", id)
	event := entity.Event{}

	err := rows.Scan(&event.EId, &event.Name, &event.Details, &event.Image, &event.City, &event.Country, &event.Place, &event.Coordinates, &event.host, &event.IsPassed, &event.Rating, &event.PostedDate, &event.price)

	if err != nil {
		return event, err
	}


	return event, nil
}

//Post gy
func (eri *EventRepoImpl) Post(event entity.Event) error {

	_, err := eri.conn.Exec("INSERT INTO events (EId, Name, Details, Image, City, Country, Place, Coordinates, host, IsPassed, Rating, PostedDate, price) values($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13)",event.EId, event.Name, event.Details, event.Image, event.City, event.Country, event.Place, event.Coordinates, event.host, event.IsPassed, event.Rating, event.PostedDate, event.price)
	if err != nil {
		return errors.New("Insertion has failed")
	}

	return nil
}


// EditEvent nn
func (eri *EventRepoImpl) EditEvent(event entity.Event) error {
	_, err := eri.conn.Exec("UPDATE events SET EId=$1, Name=$2, Details=$3, Image=$4, City=$5, Country=$6, Place=$7, Coordinates=$8, host=$9, IsPassed=$10, Rating=$11, PostedDate=$12, price=$13) WHERE id=$14",event.EId, event.Name, event.Details, event.Image, event.City, event.Country, event.Place, event.Coordinates, event.host, event.IsPassed, event.Rating, event.PostedDate, event.price)
	if err != nil {
		return errors.New("Insertion has failed")
	}

	return nil
}

// DeleteEvent delets events
func (eri *EventRepoImpl) DeleteEvent(id int) error {

_, err := cri.conn.Exec("DELETE FROM events WHERE id=$1", id)
	if err != nil {
		return errors.New("Delete has failed")
	}

	return nil
}

func (eri *EventRepoImpl) upcomingEvents() ([]entity.Event, error) {
	
	row := eri.conn.QueryRow("SELECT * FROM events WHERE IsPassed = $1", false)
	event := entity.Event{}

	err := rows.Scan(&event.EId, &event.Name, &event.Details, &event.Image, &event.City, &event.Country, &event.Place, &event.Coordinates, &event.host, &event.IsPassed, &event.Rating, &event.PostedDate, &event.price)

	if err != nil {
		return event, err
	}
}

func (eri *EventRepoImpl) rate(id, rating int) {

	_, err := cri.conn.Exec("UPDATE events SET rating=$1 WHERE EId=$2", rating, EId)
	if err != nil {
		return errors.New("rating has failed")
	}

	return nil
}


