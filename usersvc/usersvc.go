package usersvc

import "github.com/gin-gonic/gin"

// New initialize a new usersvc
func New(r *gin.Engine, on bool) {
	evt := NewEvent()
	rep := NewRepository()
	svc := NewService(rep, evt)
	ctl := NewController(svc)
	ctl.Setup(r, on)
}
