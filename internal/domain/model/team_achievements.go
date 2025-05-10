package model

import "time"

type TeamAchievement struct {
	ID         int64       `boil:"id" json:"id" toml:"id" yaml:"id"`
	TeamID     int64       `boil:"team_id" json:"team_id" toml:"team_id" yaml:"team_id"`
	PrizeID    int64       `boil:"prize_id" json:"prize_id" toml:"prize_id" yaml:"prize_id"`
	CreatedAt  time.Time   `boil:"created_at" json:"created_at" toml:"created_at" yaml:"created_at"`
	RecordHash []byte  `boil:"record_hash" json:"record_hash,omitempty" toml:"record_hash" yaml:"record_hash,omitempty"`
	TXHash     *string `boil:"tx_hash" json:"tx_hash,omitempty" toml:"tx_hash" yaml:"tx_hash,omitempty"`
}