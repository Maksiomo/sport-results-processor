package model

import "time"

type TeamAchievement struct {
	ID         int64     `json:"-"`
	TeamID     int64     `json:"team_id"`
	PrizeID    int64     `json:"prize_id"`
	CreatedAt  time.Time `json:"-"`
	TXHash     *string   `json:"-"`
}
