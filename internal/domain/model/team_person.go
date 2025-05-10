package model

import "time"

type TeamPerson struct {
	ID         int64      `boil:"id" json:"id" toml:"id" yaml:"id"`
	TeamID     int64      `boil:"team_id" json:"team_id" toml:"team_id" yaml:"team_id"`
	PersonID   int64      `boil:"person_id" json:"person_id" toml:"person_id" yaml:"person_id"`
	RoleID     int64      `boil:"role_id" json:"role_id" toml:"role_id" yaml:"role_id"`
	JoinedAt   time.Time  `boil:"joined_at" json:"joined_at" toml:"joined_at" yaml:"joined_at"`
	LeftAt     *time.Time `boil:"left_at" json:"left_at,omitempty" toml:"left_at" yaml:"left_at,omitempty"`
	CreatedAt  time.Time  `boil:"created_at" json:"created_at" toml:"created_at" yaml:"created_at"`
	RecordHash []byte     `boil:"record_hash" json:"record_hash,omitempty" toml:"record_hash" yaml:"record_hash,omitempty"`
	TXHash     *string    `boil:"tx_hash" json:"tx_hash,omitempty" toml:"tx_hash" yaml:"tx_hash,omitempty"`
}
