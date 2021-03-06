package usersvc

import "github.com/alextanhongpin/go-gin-starter/model"

// RepositoryImpl represents the implementation of the repository for user
type RepositoryImpl struct {
	// DB *db.Client
}

// GetUser takes a name and returns a user or error
func (r *RepositoryImpl) GetUser(name string) (*model.User, error) {
	return &model.User{
		Name: name,
	}, nil
}

// NewRepository returns a new pointer to the repository implementation
func NewRepository() *RepositoryImpl {
	return &RepositoryImpl{}
}
