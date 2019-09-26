package main

import (
	"github.com/gin-gonic/gin"
	"github.com/larlandis/ddd-example/pkg/user"
)

func main() {
	// Init server
	s := gin.Default()

	// GET /ddd-example/:userID handler
	s.GET("/ddd-example/:userID", user.HandlerGetUser)

	// Run server
	s.Run()
}
