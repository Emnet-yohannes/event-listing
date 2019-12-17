package entity

type Comment struct {
	CId   int
	Body  string
	owner User.UId

	eventId  Event.EId
	Official bool

	// CommentdOn Comment.CId  we could use this but it will be a blogging system
}

type FavoriteEvent struct {
	UId     int
	EventId int
}

type Ratings struct {
	UId    User.UId
	EId    Event.EId
	rating int
}
