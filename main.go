package main

import (
	"github.com/alextanhongpin/gin-starter/usersvc"
	"github.com/gin-gonic/gin"
)

func makeRouter() *gin.Engine {
	r := gin.Default()

	// Setup middlewares, logger etc
	// r.Use(cors)
	// r.Use(logger)
	// r.Use(secure)

	return r
}

func main() {
	// Setup dependencies
	// cfg := config.New()
	// db := database.New()
	r := makeRouter()

	// Setup services
	usersvc.Make(r)
	// svc1.Make(r, ...options)
	// svc2.Make(r, ...options)
	// svc3.Make(r, ...options)

	r.Run(":3000")
}
