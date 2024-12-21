package seed

import (
	"encoding/json"
	"log"
	"os"
	"realtime-score/internal/models"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

func SeedingMatch(db *gorm.DB) error {
	file, err := os.Open("internal/migration/data/match.json")
	if err != nil {
		log.Fatalf("Error opening seed data file: %v", err)
		return err
	}
	defer file.Close()

	var matches []struct {
		TeamA     string `json:"team_a"`
		ScoreA    int64  `json:"score_a"`
		TeamB     string `json:"team_b"`
		ScoreB    int64  `json:"score_b"`
		StartDate string `json:"start_date"`
	}

	if err := json.NewDecoder(file).Decode(&matches); err != nil {
		log.Fatalf("Error decoding seed data: %v", err)
		return err
	}

	for _, m := range matches {
		startDate, err := time.Parse(time.RFC3339, m.StartDate)
		if err != nil {
			log.Printf("Error parsing StartDate for match %s vs %s: %v", m.TeamA, m.TeamB, err)
			continue
		}

		match := models.Match{
			ID:        uuid.New(),
			TeamA:     m.TeamA,
			ScoreA:    m.ScoreA,
			TeamB:     m.TeamB,
			ScoreB:    m.ScoreB,
			StartDate: startDate,
		}

		// Insert into the database
		if err := db.Create(&match).Error; err != nil {
			log.Printf("Error seeding match %s vs %s: %v", match.TeamA, match.TeamB, err)
			return err
		}
		log.Printf("Seeded match %s vs %s successfully", match.TeamA, match.TeamB)
	}

	return nil
}
