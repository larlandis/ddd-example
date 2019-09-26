package main

import (
	"github.com/gin-gonic/gin"
	"github.com/larlandis/ddd-example/pkg/contact"
)

func main() {
	// Init server
	s := gin.Default()

	// GET /ddd-example/:userID handler
	s.GET("/ddd-example/:userID", contact.HandlerGetContact)

	// Run server
	s.Run()
}
