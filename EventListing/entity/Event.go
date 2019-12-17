package entity

//Event shows event intity
type Event struct {
	Name, Details, Image              string
	EId                               int
	City, Country, Place, Coordinates string
	host                              User.UId
	category                          Category.CId

	IsPassed   bool
	Rating     int
	PostedDate date.Date
	price      float32
}

// Tags []Tag //?????? how do wi put multiple tags in one field
// IsFull     bool
// notification   []
//
// Images                 []string
