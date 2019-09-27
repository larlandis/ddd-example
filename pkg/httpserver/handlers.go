package httpserver

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// HandlerGetContact get contact info
func (s *server) handlerGetContact(c *gin.Context) {
	// Parse userID from params
	userID := c.Param("userID")

	// get contact object
	contact, err := s.contact.GetContact(userID)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	// Prepare json response
	r, err := json.Marshal(contact)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	// Set and send response to client
	c.Header("Content-Type", "application/json")
	c.String(http.StatusOK, string(r))
}

func (s *server) Init() {
	s.engine = gin.Default()
}

func (s *server) RegisterRoutes(baseURL string) {
	// GET /baseURL/:userID handler
	s.engine.GET(
		fmt.Sprintf("/%s/:userID", baseURL),
		s.handlerGetContact,
	)
}

func (s *server) Run() {
	s.engine.Run()
}
