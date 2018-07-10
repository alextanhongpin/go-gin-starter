package controller

import (
	"net/http"

	"github.com/alextanhongpin/gin-starter/service"
	"github.com/gin-gonic/gin"
)

type (
	// User represents the User Controller
	User interface {
		GetUser(c *gin.Context)
		Setup(r *gin.Engine)
	}

	userController struct {
		Service service.User
	}
)

func (u *userController) GetUser(c *gin.Context) {
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

func (u *userController) Setup(r *gin.Engine) {
	e := r.Group("/users")
	e.GET(":name", u.GetUser)
}

// MakeUser returns a pointer to the userController
func MakeUser(svc service.User) User {
	return &userController{svc}
}
