package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/alextanhongpin/gin-starter/model"
)

type (
	userService interface {
		GetUser(name string) (*model.User, error)
	}

	// UserController represents the User Controller
	UserController struct {
		Service userService
	}
)

func (u *UserController) GetUser(c *gin.Context) {
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

func (u *UserController) Setup(r *gin.Engine) {
	e := r.Group("/users")
	e.GET(":name", u.GetUser)
}

// MakeUserController returns a pointer to the userController
func MakeUserController(svc userService) *UserController {
	return &UserController{svc}
}
