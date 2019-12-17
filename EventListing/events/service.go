package events

import "github.com/birukbelay/Aprojects/eventListing/entity"

// EventRepository repository
type EventServices interface {
	Events() ([]entity.Event, error)
	Event(id int) (entity.Event, error)
	UpcomingEvents() ([]entity.Event, error)

	Post(event entity.Event) error
	UpdateEvent(event entity.Event) error
	DeleteEvent(id int) error

	rate(EventId,UserId rating int)  error
	setRating(eventId int) error //this is done in the back

	getMyRating(UId, EventId int)    int

	addTag(id ...int)                           error//?? how do we add multiple tags
	notify(eventId int, tagsId ...int)          error

}


type UserServices interface {
	SignUp(user enity.User) error
	Login(username, password string) error
	Logout() error
	DeleteAcount(id int) error
	EditAcount(user enity.User) error

	MyEvents(id int)          ([]entity.Event, error)

	// commentOnCommendt(comentId string, comment entity.Comment)   error
	comment(eventId string, comment entity.Comment)     error
	
	// ask(eventId int, question Question)                error// ?? how does a user can answer  questions directly asked for him
	// answer(questionId int, answer entity.Answer)       error



}