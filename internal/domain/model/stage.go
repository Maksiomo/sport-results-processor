package model

import "time"

type Stage struct {
	ID            int64      `json:"id"`
	CompetitionID int64      `json:"competition_id"`
	Name          string     `json:"name"`
	StartTime     time.Time  `json:"start_time"`
	EndTime       *time.Time `json:"end_time,omitempty"`
	CreatedAt     time.Time  `json:"created_at"`
}
