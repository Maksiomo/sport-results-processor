package model

import "time"

type Match struct {
	ID         int64     `json:"-"`
	StageID    int64     `json:"stage_id"`
	MatchTime  time.Time `json:"match_time"`
	LocationID *int64    `json:"location_id,omitempty"`
	Metadata   []byte    `json:"metadata,omitempty"`
	CreatedAt  time.Time `json:"-"`
	TXHash     *string   `json:"-"`
}
