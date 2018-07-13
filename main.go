package main

import (
	"time"

	"github.com/alextanhongpin/go-gin-starter/config"
	"github.com/alextanhongpin/go-gin-starter/usersvc"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func newRouter() *gin.Engine {
	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "OPTIONS"},
		AllowHeaders:     []string{"Origin"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	// Setup middlewares, logger etc
	// r.Use(logger)
	// r.Use(secure)

	return r
}

func main() {
	// Setup dependencies
	cfg := config.New()
	cfg.AutomaticEnv()
	// db := database.New()

	r := newRouter()

	// Setup services
	usersvc.New(r, cfg.GetBool("usersvc_on"))
	// svc1.New(r, ...options)
	// svc2.New(r, ...options)
	// svc3.New(r, ...options)

	r.Run(":3000")
}
