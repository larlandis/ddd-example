package contact

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
)

// HandlerGetContact get contact info
func HandlerGetContact(c *gin.Context) {
	// Parse userID from params
	userID := c.Param("userID")

	// get user object
	contact, err := getContact(c, userID)
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
