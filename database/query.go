package database

import (
	"github.com/go-pg/pg/v10"

	"github.com/cynthiawilliamsa/meetmeup/graph/shared"
)

type MeetMeUpRepo struct {
	DB *pg.DB
}

func (m *MeetMeUpRepo) GetMeetups() ([]*shared.Meetup, error) {
	var meetups []*shared.Meetup
	err := m.DB.Model(&meetups).Select()
	if err != nil {
		return nil, err
	}
	return meetups, nil
}

func (m *MeetMeUpRepo) FindByID(id string) (*shared.User, error) {
	var user shared.User
	err := m.DB.Model(&user).Where("id = ?", id).First()
	if err != nil {
		return nil, err
	}
	return &user, nil
}
