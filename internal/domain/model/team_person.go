package model

import "time"

type TeamPerson struct {
	ID         int64      `json:"id"`
	TeamID     int64      `json:"team_id"`
	PersonID   int64      `json:"person_id"`
	RoleID     int64      `json:"role_id"`
	JoinedAt   time.Time  `json:"joined_at"`
	LeftAt     *time.Time `json:"left_at,omitempty"`
	CreatedAt  time.Time  `json:"created_at"`
	RecordHash []byte     `json:"record_hash,omitempty"`
	TXHash     *string    `json:"tx_hash,omitempty"`
}
