package model

import "time"

type Match struct {
	ID         int64     `json:"id"`
	StageID    int64     `json:"stage_id"`
	MatchTime  time.Time `json:"match_time"`
	LocationID *int64    `json:"location_id,omitempty"`
	Metadata   []byte    `json:"metadata,omitempty"`
	CreatedAt  time.Time `json:"created_at"`
	RecordHash []byte    `json:"record_hash,omitempty"`
	TXHash     *string   `json:"tx_hash,omitempty"`
}
