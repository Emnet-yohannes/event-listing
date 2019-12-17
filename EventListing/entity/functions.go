package entity



type EventFunctions interface {
	events() ([]entity.Event, error)
	upcomingEvents() ([]entity.Event, error)
	event(id int) (entity.Event, error)
	post(event entity.Event) error
	editEvent(event entity.Event) error
	deleteEvent(id int) error

	rate(id, rating int) error

	addTag(...id int)                           error//?? how do we add multiple tags

	notify(eventId int, ...tagsId int)          error

	participants(id int)([]User)

}

type UserFunctions interface {
	SignUp(user enity.User) error
	Login(username, password string) error
	Logout() error
	DeleteAcount(id int) error
	EditAcount(user enity.User) error

	MyEvents(id int)          ([]entity.Event, error)
	commentOnCommendt(comentId string, comment entity.Comment)   error
	commentOnEvent(eventId string, comment entity.Comment)     error
	ask(eventId int, question Question)                error// ?? how does a user can answer  questions directly asked for him
	answer(questionId int, answer entity.Answer)       error

	post(event entity.Event)                    error
	editEvent(event entity.Event)               error
	deleteEvent(id int)                         error

}
