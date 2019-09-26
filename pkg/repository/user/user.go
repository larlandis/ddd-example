package user

import (
	"encoding/json"
	"fmt"

	"github.com/larlandis/ddd-example/pkg/restclient"
)

// ByID returns user info by userID
func ByID(userID string) (*user, error) {

	// Use new restclient
	rest := restclient.NewRestClient("http://api.internal.ml.com")
	usersURL := fmt.Sprintf("/users/%s", userID)

	// Do get request from users
	bytes, err := rest.DoGet(usersURL)
	if err != nil {
		return nil, err
	}

	// Make object from response
	userFromAPI := &user{}
	if err := json.Unmarshal(bytes, &userFromAPI); err != nil {
		return nil, err
	}
	return userFromAPI, nil
}
