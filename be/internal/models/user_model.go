package models

import "github.com/google/uuid"

type User struct {
	ID         uuid.UUID `gorm:"primaryKey;column:id;type:uuid;default:uuid_generate_v4()" json:"id"`
	Username   string    `json:"username"`
	Email      string    `json:"email"`
	Password   string    `json:"password"`
	IsVerified bool      `json:"is_verified"`
	Role       string    `json:"role"`
}
