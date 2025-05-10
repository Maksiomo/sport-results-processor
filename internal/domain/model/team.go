package model

import "time"

type Team struct {
	ID             int64     `json:"id"`
	Name           string    `json:"name"`
	CountryID      int64     `json:"country_id"`
	FoundationDate time.Time `json:"foundation_date"`
	CreatedAt      time.Time `json:"created_at"`
	RecordHash     []byte    `json:"record_hash,omitempty"`
	TXHash         *string   `json:"tx_hash,omitempty"`
}
