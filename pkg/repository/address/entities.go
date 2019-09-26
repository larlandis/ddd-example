package address

type location struct {
	Name string `json:"name,omitempty"`
}

type address struct {
	Street  string   `json:"address_line"`
	City    location `json:"city"`
	Country location `json:"country"`
}
