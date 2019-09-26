package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
)

type user struct {
	UserID    string   `json:"user_id"`
	UserName  string   `json:"nickname"`
	Email     string   `json:"email"`
	Addresses []string `json:"address,omitempty"`
}

// Reponse from users api
type userApiResponse struct {
	Email    string `json:"email"`
	NickName string `json:"nickname"`
}

// Response from address api
type addressApiResponse []Address

// Address struct
type Address struct {
	Street  string   `json:"address_line"`
	City    location `json:"city"`
	Country location `json:"country"`
}

type location struct {
	Name string `json:"name,omitempty"`
}

func main() {
	// Init server
	s := gin.Default()

	// GET /ddd-example/:userID handler
	s.GET("/ddd-example/:userID", func(c *gin.Context) {
		// Parse userID from params
		userID := c.Param("userID")

		// Get user mail and nickname
		userFromAPI, err := getUser(userID)
		if err != nil {
			c.AbortWithError(http.StatusInternalServerError, err)
			return
		}

		// Get addresses
		addressesFromAPI, err := getAddresses(userID)
		if err != nil {
			c.AbortWithError(http.StatusInternalServerError, err)
			return
		}
		var addresses []string
		for _, ad := range addressesFromAPI {
			addresses = append(addresses, fmt.Sprintf(
				"%s, %s, %s",
				ad.Street,
				ad.City.Name,
				ad.Country.Name,
			))
		}

		// Make user object
		user := &user{
			UserID:    userID,
			UserName:  userFromAPI.NickName,
			Email:     userFromAPI.Email,
			Addresses: addresses,
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
	})

	// Run server
	s.Run()
}

func getUser(userID string) (*userApiResponse, error) {

	// Init request client
	client := &http.Client{}
	url := "http://api.internal.ml.com/users/%s"

	// Prepare GET request to users API
	req, err := http.NewRequest("GET", fmt.Sprintf(url, userID), nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("X-Caller-Scopes", "admin")

	// Do request
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	// Close body on finish
	defer resp.Body.Close()

	// Read response to bytes
	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	// Make object from response
	userFromAPI := &userApiResponse{}
	if err := json.Unmarshal(bytes, &userFromAPI); err != nil {
		return nil, err
	}
	return userFromAPI, nil
}

func getAddresses(userID string) (addressApiResponse, error) {

	var addressesFromAPI addressApiResponse
	// Init request client
	client := &http.Client{}
	url := "http://api.internal.ml.com/users/%s/addresses"

	// Prepare GET request to users API
	req, err := http.NewRequest("GET", fmt.Sprintf(url, userID), nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("X-Caller-Scopes", "admin")

	// Do request
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	// Close body on finish
	defer resp.Body.Close()

	// Read response to bytes
	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	// Make object from response
	if err := json.Unmarshal(bytes, &addressesFromAPI); err != nil {
		return nil, err
	}
	return addressesFromAPI, nil
}
