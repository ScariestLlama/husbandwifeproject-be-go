package models

// CardRequest is the model for card data.
type CardRequest struct {
	ID      string `datastore:"id,omitempty" json:"id,omitempty"`
	Welsh   string `datastore:"welsh" json:"welsh"`
	English string `datastore:"english" json:"english"`
}
