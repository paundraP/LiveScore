package seed

import (
	"encoding/json"
	"log"
	"os"
	"realtime-score/internal/models"
	"realtime-score/internal/pkg"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

func SeedingUser(db *gorm.DB) error {
	file, err := os.Open("internal/migration/data/users.json")
	if err != nil {
		log.Fatalf("Error opening seed data file: %v", err)
	}

	defer file.Close()

	var users []models.User
	if err := json.NewDecoder(file).Decode(&users); err != nil {
		log.Fatalf("Error decoding seed data: %v", err)
		return err
	}
	for i, user := range users {
		var existingUser models.User
		db.Where("email = ?", user.Email).First(&existingUser)
		if existingUser.ID != uuid.Nil {
			log.Printf("Skipping user %d: email %s already exists", i, user.Email)
			continue
		}

		hashedPassword, err := pkg.HashPassword(user.Password)
		if err != nil {
			log.Fatalf("Error hashing password for user %d: %v", i, err)
			return err
		}
		user.Password = hashedPassword

		result := db.Create(&user)
		if result.Error != nil {
			log.Printf("Error inserting user %d: %v", i, result.Error)
		} else {
			log.Printf("Inserted user %d: %s", i, user.Email)
		}
	}
	log.Println("Database seeding completed successfully!")
	return nil
}
