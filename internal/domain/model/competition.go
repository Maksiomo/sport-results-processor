package model

import "time"

type Competition struct {
	ID         int64     `json:"id"`
	Name       string    `json:"name"`
	SportID    int64     `json:"sport_id"`
	LocationID int64     `json:"location_id"`
	LevelID    int64     `json:"level_id"`
	CreatedAt  time.Time `json:"created_at"`
	RecordHash []byte    `json:"record_hash,omitempty"`
	TXHash     *string   `json:"tx_hash,omitempty"`
}
