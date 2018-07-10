package service

import "github.com/alextanhongpin/gin-starter/model"

type (
	// User represents the User Service
	User interface {
		GetUser(name string) (*model.User, error)
	}

	userService struct{}
)

func (svc *userService) GetUser(name string) (*model.User, error) {
	u := model.User{
		Name: name,
	}
	return &u, nil
}

// MakeUser returns a pointer to the userService
func MakeUser() User {
	return &userService{}
}
