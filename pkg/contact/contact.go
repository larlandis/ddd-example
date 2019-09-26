package contact

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/larlandis/ddd-example/pkg/repository/address"
)

// UserRepo defines behaviour of user repository
type UserRepo interface {
	ByID(userID string) (*User, error)
}

type contactService struct {
	userRepo UserRepo
}

func (serv contactService) getContact(c *gin.Context, userID string) (*Contact, error) {
	// Get mail and nickname
	user, err := serv.userRepo.ByID(userID)
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
		UserName:  user.Name,
		Email:     user.Email,
		Addresses: addresses,
	}

	return contact, nil
}

func newContactService(uRepo UserRepo) *contactService {
	return &contactService{
		userRepo: uRepo,
	}
}
