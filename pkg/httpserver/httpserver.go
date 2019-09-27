package httpserver

import (
	"github.com/gin-gonic/gin"
	"github.com/larlandis/ddd-example/pkg/contact"
)

type Server interface {
	Init()
	RegisterRoutes(baseURL string)
	Run()
}

type server struct {
	contact contact.Service
	engine  *gin.Engine // I don't like gin
}

// New returns gin server
func New(cServ contact.Service) Server {
	return &server{
		contact: cServ,
	}
}
