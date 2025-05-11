package model

import "time"

type TeamAchievement struct {
	ID         int64     `json:"id"`
	TeamID     int64     `json:"team_id"`
	PrizeID    int64     `json:"prize_id"`
	CreatedAt  time.Time `json:"created_at"`
	RecordHash []byte    `json:"record_hash,omitempty"`
	TXHash     *string   `json:"tx_hash,omitempty"`
}
