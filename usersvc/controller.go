package usersvc

import (
	"net/http"

	"github.com/alextanhongpin/go-gin-starter/model"
	"github.com/gin-gonic/gin"
)

type (
	service interface {
		GetUser(name string) (*model.User, error)
		// PostUser(user model.User) error
		// DeleteUser(name string) error
		// UpdateUser(user model.User) error
	}

	// Controller represents the User Controller
	Controller struct {
		Service service
	}
)

// GetUser represents the controller to get the users
func (u *Controller) GetUser(c *gin.Context) {
	name := c.Param("name")
	user, err := u.Service.GetUser(name)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, &user)
}

// Setup binds the router to the user controllers
func (u *Controller) Setup(r *gin.Engine) {
	e := r.Group("/users")
	{
		e.GET("/:name", u.GetUser)
		// e.POST("", u.PostUser)
		// e.DELETE("", u.DeleteUser)
		// e.Update("", u.UpdateUser)
	}
}

// NewController returns a pointer to the userController
func NewController(s service) *Controller {
	return &Controller{s}
}
