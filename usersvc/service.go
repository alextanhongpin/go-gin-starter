package usersvc

import "github.com/alextanhongpin/go-gin-starter/model"

type (
	repository interface {
		GetUser(name string) (*model.User, error)
	}

	event interface {
	}

	// Service represents the business layer for the user service
	Service struct {
		Repository repository
		Event      event
	}
)

// GetUser returns the user by name
func (svc *Service) GetUser(name string) (*model.User, error) {
	u, err := svc.Repository.GetUser(name)
	if err != nil {
		return nil, err
	}
	// Example of emitting webhook events
	// svc.Event.CreatedUser(name)
	return u, nil
}

// NewService returns a pointer to the userService
func NewService(r repository, e event) *Service {
	return &Service{r, e}
}
