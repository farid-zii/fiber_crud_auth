package main

import (
	"fiber-crud-auth/controllers"
	"fiber-crud-auth/databases"
	"fiber-crud-auth/repositories"
	"fiber-crud-auth/routes"
	"fiber-crud-auth/seeder"
	"fiber-crud-auth/services"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil{
		log.Fatal("Error loading .env file")
	}
	app := fiber.New()

	//inisialisasi db
	databases.InitDatabase()

	// Menjalankan seeder
    seeder.SeedAll(databases.DB)

	//inisialisasi repository
	userRepo := repositories.NewUserRepository()
	bookRepo := repositories.NewBookRepository()

	// inisialisasi services
	authService := services.NewAuthService(userRepo,os.Getenv("JWT_SECRET"))
	bookService := services.NewBookService(bookRepo)


	//inisialisasi Controllers
	authController := controllers.NewAuthController(authService)
	bookContoller := controllers.NewBookController(bookService)


	//Setup routes
	routes.SetupRoutes(app, authController, bookContoller)

	//status file servcer untuk upload folder
	app.Static("/uploads","./public/uploads")

	app.Listen(":3000")
}