package usersvc

import "github.com/gin-gonic/gin"

// Make initialize a new usersvc
func Make(r *gin.Engine) {
	rep := MakeRepository()
	svc := MakeService(rep)
	ctl := MakeController(svc)
	ctl.Setup(r)
}
