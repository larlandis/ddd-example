package main

import (
	"github.com/larlandis/ddd-example/pkg/contact"
	"github.com/larlandis/ddd-example/pkg/httpserver"
	"github.com/larlandis/ddd-example/pkg/repository/address"
	"github.com/larlandis/ddd-example/pkg/repository/user"
)

func main() {

	// Init repositories
	userRepo := user.NewRepo()
	// userRepo := user.NewMockRepo()
	addressRepo := address.NewRepo()

	// Init services
	contactServ := contact.NewService(userRepo, addressRepo)

	server := httpserver.New(contactServ)
	server.Init()
	server.RegisterRoutes("ddd-example")
	server.Run()

}
