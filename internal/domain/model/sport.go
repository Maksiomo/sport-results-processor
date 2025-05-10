package model

import "time"

type Sport struct {
	ID          int64     `json:"id"`
	Name        string    `json:"name"`
	MinTeamSize int       `json:"min_team_size"`
	MaxTeamSize int       `json:"max_team_size"`
	Description *string   `json:"description,omitempty"`
	CreatedAt   time.Time `json:"created_at"`
	RecordHash  []byte    `json:"record_hash,omitempty"`
	TXHash      *string   `json:"tx_hash,omitempty"`
}
