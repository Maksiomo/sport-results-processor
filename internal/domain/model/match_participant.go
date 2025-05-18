package model

import "time"

type MatchParticipant struct {
	ID         int64     `json:"-"`
	MatchID    int64     `json:"match_id"`
	TeamID     int64     `json:"team_id"`
	Score      int       `json:"score"`
	IsWinner   bool      `json:"is_winner"`
	CreatedAt  time.Time `json:"-"`
	TXHash     *string   `json:"-"`
}
