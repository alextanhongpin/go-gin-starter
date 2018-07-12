package usersvc

import "github.com/alextanhongpin/go-gin-starter/model"

type repository interface {
	GetUser(name string) (*model.User, error)
}

// Service represents the business layer for the user service
type Service struct {
	Repository repository
}

// GetUser returns the user by name
func (svc *Service) GetUser(name string) (*model.User, error) {
	u, err := svc.Repository.GetUser(name)
	if err != nil {
		return nil, err
	}
	return u, nil
}

// MakeService returns a pointer to the userService
func MakeService(r repository) *Service {
	return &Service{r}
}
