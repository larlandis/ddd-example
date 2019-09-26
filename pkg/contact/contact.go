package contact

import (
	"github.com/gin-gonic/gin"
)

// UserRepo defines behaviour of user repository
type UserRepo interface {
	ByID(userID string) (*User, error)
}

// AddressRepo defines behaviour of address repository
type AddressRepo interface {
	ByID(userID string) ([]string, error)
}

type contactService struct {
	userRepo    UserRepo
	addressRepo AddressRepo
}

func (serv contactService) getContact(c *gin.Context, userID string) (*Contact, error) {
	// Get mail and nickname
	user, err := serv.userRepo.ByID(userID)
	if err != nil {
		return nil, err
	}

	// Get addresses
	addresses, err := serv.addressRepo.ByID(userID)
	if err != nil {
		return nil, err
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

func newContactService(uRepo UserRepo, aRepo AddressRepo) *contactService {
	return &contactService{
		userRepo:    uRepo,
		addressRepo: aRepo,
	}
}
