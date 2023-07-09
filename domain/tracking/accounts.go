package tracking

type Account struct {
	Created    string `json:"created"`
	ExternalID string `json:"external_id"`
	FirstName  string `json:"first_name"`
	LastName   string `json:"last_name"`
	IsActive   bool   `json:"is_active"`
}
