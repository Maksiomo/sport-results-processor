package model

import "time"

type Team struct {
	ID             int64     `json:"-"`
	Name           string    `json:"name"`
	CountryID      int64     `json:"country_id"`
	FoundationDate time.Time `json:"foundation_date"`
	CreatedAt      time.Time `json:"-"`
	TXHash         *string   `json:"-"`
}
