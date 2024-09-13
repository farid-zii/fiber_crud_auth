package controllers

import (
	"fiber-crud-auth/models"
	"fiber-crud-auth/services"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
)

type BookController struct {
	service services.BookService
}

func NewBookController(service services.BookService) *BookController{
	return &BookController{service}
}

//Create Book utk membuat buku baru dengan upload file
func (c *BookController) CreateBook(ctx *fiber.Ctx)error{
	book := new(models.Book)

	if err:=ctx.BodyParser(book); err != nil{
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error":err.Error()})
	}

	file,err := ctx.FormFile("image")
	if err == nil{
		filename := strconv.Itoa(int(time.Now().Unix()))+"_"+file.Filename
		filePath := "./public/uploads/"+filename
		if err := ctx.SaveFile(file,filePath); err != nil{
			return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error":err.Error()})
		}
		book.Image = filename
	} 

	if err := c.service.CreateBook(book);err != nil{
			return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error":err.Error()})
	}

	return ctx.Status(fiber.StatusCreated).JSON(book)
}

//GetAllBooks untuk mendapatkab semua buku
func(c *BookController) GetAllBooks(ctx *fiber.Ctx)error{
	books,err:= c.service.GetAllBooks()
	if err != nil{
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error":err.Error()})
	}

	return ctx.JSON(books)
}

//GETBOOKByID untuk mendapayakna buku berdasarkan ID
func(c *BookController) GetBookByID(ctx *fiber.Ctx)error{
	id,err:= strconv.Atoi(ctx.Params("id"))
	if err != nil{
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error":"invalid book ID"})
	}

	book,err :=c.service.GetBookByID(uint(id))
	if err != nil{
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{"error":"Book Not Found"})
	}

	return ctx.JSON(book)
}

func (c *BookController) UpdateBook(ctx *fiber.Ctx)error{
	id,err := strconv.Atoi(ctx.Params("id"))
	if err!=nil{
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error":"Invalid book ID"})
	}

	book,err := c.service.GetBookByID(uint(id))
	if err !=nil{
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{"error":"Book Not Found"})
	}

	if err := ctx.BodyParser(book); err !=nil{
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error":err.Error()})
	}

	file,err := ctx.FormFile("image")
	if err == nil{
		filename := strconv.Itoa(int(time.Now().Unix()))+"_"+file.Filename
		filepath := "./public/uploads/"+filename
		if err := ctx.SaveFile(file, filepath); err !=nil{
			return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error":err.Error()})
		}
		book.Image=filename
	}

	if err:=c.service.UpdateBook(book);err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error":err.Error()})
	}

	return ctx.JSON(book)
}

//DELETE 
func (c *BookController) DeleteBook(ctx *fiber.Ctx)error{
	id,err := strconv.Atoi(ctx.Params("id"))
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error":"Invalid book ID"})
	}

	if err := c.service.DeleteBook(uint(id)); err != nil{
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error":err.Error()})
	}

	return ctx.SendStatus(fiber.StatusNoContent)
}