package model

import "time"

type Person struct {
	ID         int64     `json:"-"`
	Name       string    `json:"name"`
	BirthDate  time.Time `json:"birth_date"`
	CountryID  int64     `json:"country_id"`
	CreatedAt  time.Time `json:"-"`
	RecordHash string    `json:"-"`
	TXHash     *string   `json:"-"`
}
