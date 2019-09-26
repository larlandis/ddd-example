package user

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
)

// HandlerGetUser get user info
func HandlerGetUser(c *gin.Context) {
	// Parse userID from params
	userID := c.Param("userID")

	// get user object
	user, err := getUser(c, userID)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	// Prepare json response
	r, err := json.Marshal(user)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	// Set and send response to client
	c.Header("Content-Type", "application/json")
	c.String(http.StatusOK, string(r))
}
