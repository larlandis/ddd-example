package address

import (
	"encoding/json"
	"fmt"

	"github.com/larlandis/ddd-example/pkg/contact"
	"github.com/larlandis/ddd-example/pkg/restclient"
)

type addressRepo struct{}

// NewRepo creates and returns new user repository
func NewRepo() contact.AddressRepo {
	return &addressRepo{}
}

// ByID returns array of address from a user
func (*addressRepo) ByID(userID string) ([]string, error) {

	// Use new restclient
	rest := restclient.NewRestClient("http://api.internal.ml.com")
	url := fmt.Sprintf("/users/%s/addresses/", userID)

	// Do get request from users
	bytes, err := rest.DoGet(url)
	if err != nil {
		return nil, err
	}

	// Make object from response
	addressesFromAPI := []address{}
	if err := json.Unmarshal(bytes, &addressesFromAPI); err != nil {
		return nil, err
	}

	// Parse data from api
	var addresses []string
	for _, ad := range addressesFromAPI {
		addresses = append(addresses, fmt.Sprintf(
			"%s, %s, %s",
			ad.Street,
			ad.City.Name,
			ad.Country.Name,
		))
	}

	return addresses, nil
}
