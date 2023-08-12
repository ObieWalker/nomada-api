package main

import (
	"fmt"
	"github.com/obiewalker/nomada-api/pkg/database"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/obiewalker/nomada-api/config"
)

// func (r *Repository) CreateBook(context *fiber.Ctx) error {
// 	return context.SendString("Obi app")
// 	book := Book{}

// 	err := context.BodyParser(&book)

// 	if err != nil {

// 		context.Status(http.StatusUnprocessableEntity).JSON(
// 			&fiber.Map{"message": "request failed"})
// 		return err
// 	}

// 	err = r.DB.Create(&book).Error
// 	if err != nil {
// 		context.Status(http.StatusBadRequest).JSON(
// 			&fiber.Map{"message": "could not create book"})
// 		return err
// 	}

// 	context.Status(http.StatusOK).JSON(&fiber.Map{
// 		"message": "book has been added"})
// 	return nil
// }

// func (r *Repository) DeleteBook(context *fiber.Ctx) error {
// 	bookModel := models.Books{}
// 	id := context.Params("id")
// 	if id == "" {
// 		context.Status(http.StatusInternalServerError).JSON(&fiber.Map{
// 			"message": "id cannot be empty",
// 		})
// 		return nil
// 	}

// 	err := r.DB.Delete(bookModel, id)

// 	if err.Error != nil {
// 		context.Status(http.StatusBadRequest).JSON(&fiber.Map{
// 			"message": "could not delete book",
// 		})
// 		return err.Error
// 	}
// 	context.Status(http.StatusOK).JSON(&fiber.Map{
// 		"message": "book delete successfully",
// 	})
// 	return nil
// }

// func (r *Repository) GetBooks(context *fiber.Ctx) error {
// 	bookModels := &[]models.Books{}

// 	err := r.DB.Find(bookModels).Error
// 	if err != nil {
// 		context.Status(http.StatusBadRequest).JSON(
// 			&fiber.Map{"message": "could not get books"})
// 		return err
// 	}

// 	context.Status(http.StatusOK).JSON(&fiber.Map{
// 		"message": "books fetched successfully",
// 		"data":    bookModels,
// 	})
// 	return nil
// }

// func (r *Repository) GetBookByID(context *fiber.Ctx) error {

// 	id := context.Params("id")
// 	bookModel := &models.Books{}
// 	if id == "" {
// 		context.Status(http.StatusInternalServerError).JSON(&fiber.Map{
// 			"message": "id cannot be empty",
// 		})
// 		return nil
// 	}

// 	fmt.Println("the ID is", id)

// 	err := r.DB.Where("id = ?", id).First(bookModel).Error
// 	if err != nil {
// 		context.Status(http.StatusBadRequest).JSON(
// 			&fiber.Map{"message": "could not get the book"})
// 		return err
// 	}
// 	context.Status(http.StatusOK).JSON(&fiber.Map{
// 		"message": "book id fetched successfully",
// 		"data":    bookModel,
// 	})
// 	return nil
// }


func main() {

	fmt.Println("this has started")

	database.ConnectDb()
	
	app := fiber.New()
	app.Use(cors.New(cors.Config{
    AllowOrigins:     config.GetEnvStr("CLIENT_ORIGIN"),
		AllowHeaders:     "Origin, Content-Type, Accept",
    AllowCredentials: true,
	}))
	setupRoutes(app)
	app.Listen(":3000")
}

