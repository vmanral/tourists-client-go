package tourists

// Order -
type Order struct {
	ID    int         `json:"id,omitempty"`
	Items []OrderItem `json:"data,omitempty"`
}

// OrderItem -
type OrderItem struct {
	Tourist   Tourist `json:"tourist"`
	Quantity  int    `json:"quantity"`
}

// Tourist -
type Tourist struct {
	ID                int            `json:"id"`
	Tourist_Name      string         `json:"tourist_name"`
	Tourist_Email     string         `json:"tourist_email"`
	Tourist_Location  string         `json:"tourist_location"`
	Createdat         string         `json:"Createdat"`
}
