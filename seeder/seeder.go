package seeder

import (
	"fiber-crud-auth/models"
	"log"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// SeedUsers untuk menambahkan data user
func SeedUsers(db *gorm.DB) {
    password, _ := bcrypt.GenerateFromPassword([]byte("password"), bcrypt.DefaultCost)

    users := []models.User{
        {
            Name:     "Admin",
            Email:    "admin@example.com",
            Password: string(password),
        },
        {
            Name:     "User",
            Email:    "user@example.com",
            Password: string(password),
        },
    }

    for _, user := range users {
        err := db.Where(models.User{Email: user.Email}).FirstOrCreate(&user).Error
        if err != nil {
            log.Fatalf("cannot seed users table: %v", err)
        }
    }
}

// SeedBooks untuk menambahkan data book
func SeedBooks(db *gorm.DB) {
    books := []models.Book{
        {
            Title:       "Go Programming",
            Description: "Learn Go programming from scratch",
            Image:       "go_programming.png",
        },
        {
            Title:       "Learning Fiber",
            Description: "Master the Fiber web framework in Go",
            Image:       "learning_fiber.png",
        },
    }

    for _, book := range books {
        err := db.Where(models.Book{Title: book.Title}).FirstOrCreate(&book).Error
        if err != nil {
            log.Fatalf("cannot seed books table: %v", err)
        }
    }
}

// SeedAll untuk menjalankan semua seeder
func SeedAll(db *gorm.DB) {
    SeedUsers(db)
    SeedBooks(db)
}
