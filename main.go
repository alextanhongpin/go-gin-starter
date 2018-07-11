package main

import (
	"github.com/alextanhongpin/gin-starter/controller"
	"github.com/alextanhongpin/gin-starter/service"

	"github.com/gin-gonic/gin"
)

func makeRouter() *gin.Engine {
	r := gin.Default()

	// Setup middlewares, logger etc

	return r
}

func makeUserRoute(r *gin.Engine) {
	svc := service.MakeUserService()
	ctrl := controller.MakeUserController(svc)

	ctrl.Setup(r)
}

func main() {
	r := makeRouter()

	makeUserRoute(r)

	r.Run(":3000")
}
