package main

import (
	"github.com/gin-gonic/gin"
	"github.com/larlandis/ddd-example/pkg/contact"
	"github.com/larlandis/ddd-example/pkg/repository/address"
	"github.com/larlandis/ddd-example/pkg/repository/user"
)

func main() {
	// Init server
	s := gin.Default()

	// Init repositories
	userRepo := user.NewRepo()
	// userRepo := user.NewMockRepo()
	addressRepo := address.NewRepo()

	// GET /ddd-example/:userID handler
	s.GET("/ddd-example/:userID", contact.HandlerGetContact(userRepo, addressRepo))

	// Run server
	s.Run()
}
