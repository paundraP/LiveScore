package services

import (
	"realtime-score/internal/models"
	"realtime-score/internal/repository"
)

type MatchService struct {
	MatchRepo *repository.MatchRepository
}

func NewMatchService(repo *repository.MatchRepository) *MatchService {
	return &MatchService{MatchRepo: repo}
}

func (s *MatchService) UpdateMatchScore(match_id string, scoreA, scoreB int64) (*models.Match, error) {
	match, err := s.MatchRepo.GetMatchById(match_id)
	if err != nil || match == nil {
		return &models.Match{}, err
	}

	match.ScoreA = scoreA
	match.ScoreB = scoreB

	if err := s.MatchRepo.UpdateMatchScore(match); err != nil {
		return nil, err
	}
	return match, nil
}
