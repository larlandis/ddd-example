package user

import "github.com/larlandis/ddd-example/pkg/contact"

type userMockRepo struct{}

// NewMockRepo creates and returns new user repository, with false hardcoded data
func NewMockRepo() contact.UserRepo {
	return &userMockRepo{}
}

func (userMockRepo) ByID(_ string) (*contact.User, error) {
	return &contact.User{
		Name:  "UNUSERMAS",
		Email: "mail@example.com",
	}, nil
}
