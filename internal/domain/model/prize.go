package model

import "time"

type Prize struct {
	ID            int64     `json:"-"`
	CompetitionID int64     `json:"competition_id"`
	PlaceBracket  string    `json:"place_bracket"`
	CurrencyCode  string    `json:"currency_code"`
	PrizeAmount   int64     `json:"prize_amount"`
	CreatedAt     time.Time `json:"-"`
	TXHash        *string   `json:"-"`
}
