package routes

import (
	"fiber-crud-auth/controllers"
	"fiber-crud-auth/middleware"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App, authController *controllers.AuthController, bookController *controllers.BookController){
	api := app.Group("/api")

	// Routes untuk autentikasi
	api.Post("/register",authController.Login)
	api.Post("/login",authController.Login)


	// Routes utk buku
	bookRoutes := api.Group("/books",middleware.JWTMiddleware)
	bookRoutes.Post("/",bookController.CreateBook)
	bookRoutes.Get("/",bookController.GetAllBooks)
	bookRoutes.Get("/:id",bookController.GetBookByID)
	bookRoutes.Put("/:id",bookController.UpdateBook)
	bookRoutes.Delete("/:id",bookController.DeleteBook)
}