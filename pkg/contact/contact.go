package contact

import (
	"encoding/json"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/larlandis/ddd-example/pkg/restclient"
)

func getContact(c *gin.Context, userID string) (*Contact, error) {
	// Get user mail and nickname
	userFromAPI, err := userByID(userID)
	if err != nil {
		return nil, err
	}

	// Get addresses
	addressesFromAPI, err := addressByID(userID)
	if err != nil {
		return nil, err
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

	// Make contact object
	contact := &Contact{
		UserID:    userID,
		UserName:  userFromAPI.NickName,
		Email:     userFromAPI.Email,
		Addresses: addresses,
	}

	return contact, nil
}

func userByID(userID string) (*userApiResponse, error) {

	// Use new restclient
	rest := restclient.NewRestClient("http://api.internal.ml.com")
	url := fmt.Sprintf("/users/%s", userID)

	// Do get request from users
	bytes, err := rest.DoGet(url)
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

func addressByID(userID string) (addressApiResponse, error) {

	// Use new restclient
	rest := restclient.NewRestClient("http://api.internal.ml.com")
	url := fmt.Sprintf("/users/%s/addresses/", userID)

	// Do get request from users
	bytes, err := rest.DoGet(url)
	if err != nil {
		return nil, err
	}

	// Make object from response
	addressesFromAPI := addressApiResponse{}
	if err := json.Unmarshal(bytes, &addressesFromAPI); err != nil {
		return nil, err
	}
	return addressesFromAPI, nil
}
