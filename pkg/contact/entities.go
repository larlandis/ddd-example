package contact

// Contact entity
type Contact struct {
	UserID    string   `json:"user_id"`
	UserName  string   `json:"nickname"`
	Email     string   `json:"email"`
	Addresses []string `json:"address,omitempty"`
}
