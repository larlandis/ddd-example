package address

import (
	"encoding/json"
	"fmt"

	"github.com/larlandis/ddd-example/pkg/restclient"
)

// ByID returns array of address from a user
func ByID(userID string) ([]address, error) {

	// Use new restclient
	rest := restclient.NewRestClient("http://api.internal.ml.com")
	usersURL := fmt.Sprintf("/users/%s/addresses/", userID)

	// Do get request from users
	bytes, err := rest.DoGet(usersURL)
	if err != nil {
		return nil, err
	}

	// Make object from response
	addressesFromAPI := []address{}
	if err := json.Unmarshal(bytes, &addressesFromAPI); err != nil {
		return nil, err
	}
	return addressesFromAPI, nil
}
