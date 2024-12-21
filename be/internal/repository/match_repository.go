package repository

import (
	"realtime-score/internal/models"

	"gorm.io/gorm"
)

type MatchRepository struct {
	db *gorm.DB
}

func NewMatchRepository(db *gorm.DB) *MatchRepository {
	return &MatchRepository{db: db}
}

func (r *MatchRepository) GetMatchById(match_id string) (*models.Match, error) {
	var match models.Match

	if err := r.db.First(&match, "id = ?", match_id).Error; err != nil {
		return &models.Match{}, err
	}

	return &match, nil
}

func (r *MatchRepository) UpdateMatchScore(match *models.Match) error {
	return r.db.Save(match).Error
}
