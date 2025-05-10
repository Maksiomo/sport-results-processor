package model

import "time"

type MatchParticipant struct {
	ID         int64     `json:"id"`
	MatchID    int64     `json:"match_id"`
	TeamID     int64     `json:"team_id"`
	Score      int       `json:"score"`
	IsWinner   bool      `json:"is_winner"`
	CreatedAt  time.Time `json:"created_at"`
	RecordHash []byte    `json:"record_hash,omitempty"`
	TXHash     *string   `json:"tx_hash,omitempty"`
}
