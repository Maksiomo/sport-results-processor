package model

import "time"

type Sport struct {
	ID          int64     `json:"-"`
	Name        string    `json:"name"`
	MinTeamSize int       `json:"min_team_size"`
	MaxTeamSize int       `json:"max_team_size"`
	Description *string   `json:"description,omitempty"`
	CreatedAt   time.Time `json:"-"`
	RecordHash  string    `json:"-"`
	TXHash      *string   `json:"-"`
}
