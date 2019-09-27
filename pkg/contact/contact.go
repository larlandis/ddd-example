package contact

//Service defines the interface of contact functionality
type Service interface {
	GetContact(userID string) (*Contact, error)
}

// UserRepo defines behaviour of user repository
type UserRepo interface {
	ByID(userID string) (*User, error)
}

// AddressRepo defines behaviour of address repository
type AddressRepo interface {
	ByID(userID string) ([]string, error)
}

type service struct {
	userRepo    UserRepo
	addressRepo AddressRepo
}

func (serv service) GetContact(userID string) (*Contact, error) {
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

// NewService returns contact service implementation
func NewService(uRepo UserRepo, aRepo AddressRepo) Service {
	return &service{
		userRepo:    uRepo,
		addressRepo: aRepo,
	}
}
