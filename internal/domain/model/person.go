package model

import "time"

type Person struct {
	ID         int64     `json:"id"`
	Name       string    `json:"name"`
	BirthDate  time.Time `json:"birth_date"`
	CountryID  int64     `json:"country_id"`
	CreatedAt  time.Time `json:"created_at"`
	RecordHash []byte    `json:"record_hash,omitempty"`
	TXHash     *string   `json:"tx_hash,omitempty"`
}
