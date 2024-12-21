package models

import (
	"time"

	"github.com/google/uuid"
)

type Match struct {
	ID        uuid.UUID `gorm:"primaryKey;column:id;type:uuid;default:uuid_generate_v4()" json:"id"`
	TeamA     string    `json:"team_a" binding:"required"`
	ScoreA    int64     `json:"score_a" binding:"required"`
	TeamB     string    `json:"team_b" binding:"required"`
	ScoreB    int64     `json:"score_b" binding:"required"`
	StartDate time.Time `gorm:"type:timestamp with time zone" json:"start_date"`
	CreatedAt time.Time `gorm:"type:timestamp with time zone" json:"created_at"`
}
