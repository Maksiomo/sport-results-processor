package model

import "time"

type Competition struct {
	ID         int64     `json:"-"`
	Name       string    `json:"name"`
	SportID    int64     `json:"sport_id"`
	LocationID int64     `json:"location_id"`
	LevelID    int64     `json:"level_id"`
	CreatedAt  time.Time `json:"-"`
	RecordHash string    `json:"-"`
	TXHash     *string   `json:"-"`
}
