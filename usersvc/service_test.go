package usersvc

import (
	"errors"
	"testing"

	"github.com/alextanhongpin/go-gin-starter/model"
)

type nullRepository struct{}

func (n *nullRepository) GetUser(name string) (*model.User, error) {
	if name != "john doe" {
		return nil, errors.New("user does not exist")
	}
	return &model.User{
		Name: name,
	}, nil
}

type nullEvent struct{}

func TestGetUser(t *testing.T) {
	s := NewService(&nullRepository{}, &nullEvent{})
	u, err := s.GetUser("john doe")
	if err != nil {
		t.Error("got error", err)
	}
	expected := "john doe"
	got := u.Name

	if expected != expected {
		t.Errorf("want %s, got %s", expected, got)
	}
}
