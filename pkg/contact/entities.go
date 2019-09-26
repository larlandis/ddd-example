package contact

type Contact struct {
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
