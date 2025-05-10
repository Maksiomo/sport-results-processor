package model

import "time"

type Location struct {
	ID          int64     `json:"id"`
	CountryID   int64     `json:"country_id"`
	State       *string   `json:"state,omitempty"`
	City        *string   `json:"city,omitempty"`
	Address     *string   `json:"address,omitempty"`
	FullAddress string    `json:"full_address"`
	CreatedAt   time.Time `json:"created_at"`
}
