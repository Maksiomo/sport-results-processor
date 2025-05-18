package model

import "time"

type PersonSport struct {
	ID        int64     `json:"id"`
	PersonID  int64     `json:"person_id"`
	SportID   int64     `json:"sport_id"`
	CreatedAt time.Time `json:"created_at"`
}