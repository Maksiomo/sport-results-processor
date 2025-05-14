package model

import "time"

type CompetitionTeam struct {
	ID            int64     `json:"-"`
	TeamID        int64     `json:"team_id"`
	CompetitionID int64     `json:"competition_id"`
	CreatedAt     time.Time `json:"-"`
	TXHash        *string   `json:"-"`
}
