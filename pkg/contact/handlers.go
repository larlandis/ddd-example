package contact

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
)

// HandlerGetContact get contact info
func HandlerGetContact(userRepo UserRepo) func(*gin.Context) {
	return func(c *gin.Context) {
		// Parse userID from params
		userID := c.Param("userID")

		// serv := newContactService(user.NewRepo())
		serv := newContactService(userRepo)

		// get contact object
		contact, err := serv.getContact(c, userID)
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
}
