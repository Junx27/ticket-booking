package database

import (
	"log"
	"time"

	"github.com/Junx27/ticket-booking/entity"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func SeedUsers(db *gorm.DB) {

	users := []entity.User{
		{
			Email:       "admin@example.com",
			Password:    "adminpassword",
			FirsName:    "Admin",
			LastName:    "Test",
			PhoneNumber: "1234567890",
			Role:        "admin",
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		},
		{
			Email:       "user@example.com",
			Password:    "userpassword",
			FirsName:    "Regular",
			LastName:    "Test",
			PhoneNumber: "0987654321",
			Role:        "customer",
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		},
		{
			Email:       "provider@example.com",
			Password:    "providerpassword",
			FirsName:    "Provider",
			LastName:    "Test",
			PhoneNumber: "1122334455",
			Role:        "provider",
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		},
	}
	for _, user := range users {

		passwordHash, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
		if err != nil {
			log.Fatalf("Failed to hash password: %v", err)
		}
		user.Password = string(passwordHash)

		if err := db.Create(&user).Error; err != nil {
			log.Fatalf("Failed to insert user data: %v", err)
		}
	}

	log.Println("Users seeded successfully!")

}
