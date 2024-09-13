package controllers

import (
	"fiber-crud-auth/models"
	"fiber-crud-auth/services"

	"github.com/gofiber/fiber/v2"
)

type AuthController struct {
	service services.AuthService
}

func NewAuthController(service services.AuthService) *AuthController{
	return &AuthController{service}
}

//Register Untuk Medaftarkan user BAru
func (c *AuthController) Register(ctx *fiber.Ctx) error{
	user := new(models.User)
	
	if err := ctx.BodyParser(user);err != nil{
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error":err.Error()})
	}

	if err := c.service.RegisterUser(user); err != nil{
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error":err.Error()})
	}

	return ctx.Status(fiber.StatusCreated).JSON(fiber.Map{"message":"user registered"})
}


//login Untuk mengauntetikasi user dan menghasilkan token jwt

func (c *AuthController) Login(ctx *fiber.Ctx)error{
	input :=struct{
		Email string `json:"email"`
		Password string `json:"password"`
	}{}

	if err := ctx.BodyParser(&input); err != nil{
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error":err.Error()})
	}

	token, err := c.service.LoginUser(input.Email,input.Password)
	if err!=nil{
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error":err.Error()})
	}

	return ctx.JSON(fiber.Map{"token":token})
}