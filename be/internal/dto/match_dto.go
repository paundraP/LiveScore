package dto

type UpdateScore struct {
	ScoreA int64 `json:"score_a" binding:"required"`
	ScoreB int64 `json:"score_b" binding:"required"`
}
