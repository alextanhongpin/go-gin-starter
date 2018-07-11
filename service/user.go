package service

import "github.com/alextanhongpin/gin-starter/model"

type UserService struct{}

func (svc *UserService) GetUser(name string) (*model.User, error) {
	u := model.User{
		Name: name,
	}
	return &u, nil
}

// MakeUserService returns a pointer to the userService
func MakeUserService() *UserService {
	return &UserService{}
}
