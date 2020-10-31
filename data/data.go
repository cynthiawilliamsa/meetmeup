package data

import "github.com/cynthiawilliamsa/meetmeup/graph/shared"

var Meetups = []*shared.Meetup{
	{
		ID:          "1",
		Name:        "Fun Meetup",
		Description: "A fun meetup",
		UserID:      "1",
	},
	{
		ID:          "2",
		Name:        "Cool Meetup",
		Description: "The best meetup",
		UserID:      "2",
	},
}
var Users = []*shared.User{
	{
		ID:       "1",
		Username: "Cynthia",
		Email:    "cynthia@me.com",
	},
	{
		ID:       "2",
		Username: "Joe",
		Email:    "joe@me.com",
	},
}
