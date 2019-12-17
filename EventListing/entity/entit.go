package entity

type Category struct {
	Name        string
	CId         int
	Icon        string
	Description string
}

type Tag struct {
	Ids      int
	Name     string
	category Category.CId
}

type Notification struct {
	notification Event.EId
	status       bool //to show that it is seenor not
	EndDate      date.Date
}

// // ................... optional.................

// type UsersParticipated struct { //users participated on an event

// 	UId User.UId
// 	EId Event.EId
// }
// type Question struct { //questions can only be asked on an event and con only be answered by owners of the eve3nt

// 	QId     int
// 	Body    string
// 	eventId Event.Ids
// 	owner   User.Id
// }

// type Answer struct{
// 	Ansid int
// 	Body    string
// 	Qid Question.QId
// 	owner   User.Id

// }
