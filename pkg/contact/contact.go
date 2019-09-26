package contact

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/larlandis/ddd-example/pkg/repository/address"
	"github.com/larlandis/ddd-example/pkg/repository/user"
)

func getContact(c *gin.Context, userID string) (*Contact, error) {
	// Get user mail and nickname
	userFromAPI, err := user.ByID(userID)
	if err != nil {
		return nil, err
	}

	// Get addresses
	addressesFromAPI, err := address.ByID(userID)
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
		UserName:  userFromAPI.NickName, // Acá hay un problema!
		Email:     userFromAPI.Email,
		Addresses: addresses,
	}

	return contact, nil
}
