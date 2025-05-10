package model

import "time"

type Prize struct {
	ID            int64     `json:"id"`
	CompetitionID int64     `json:"competition_id"`
	PlaceBracket  string    `json:"place_bracket"`
	CurrencyCode  string    `json:"currency_code"`
	PrizeAmount   int64     `json:"prize_amount"`
	CreatedAt     time.Time `json:"created_at"`
	RecordHash    []byte    `json:"record_hash,omitempty"`
	TXHash        *string   `json:"tx_hash,omitempty"`
}
