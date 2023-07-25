package main

import (
	"fmt"
	"net/http"
	// "os"
	// "github.com/obiewalker/nomada-api/pkg/db/models"
	"github.com/obiewalker/nomada-api/pkg/database"
	// "github.com/obiewalker/nomada-api/user"
	// "github.com/obiewalker/nomada-api/pkg/storage"
	"github.com/gofiber/fiber/v2"
	// "github.com/gofiber/fiber/middleware"
	"gorm.io/gorm"
)

type Book struct {
	Author    string `json:"author"`
	Title     string `json:"title"`
	Publisher string `json:"publisher"`
}

type Repository struct {
	DB *gorm.DB
}

func (r *Repository) CreateBook(context *fiber.Ctx) error {
	return context.SendString("Obi app")
	book := Book{}

	err := context.BodyParser(&book)

	if err != nil {

		context.Status(http.StatusUnprocessableEntity).JSON(
			&fiber.Map{"message": "request failed"})
		return err
	}

	err = r.DB.Create(&book).Error
	if err != nil {
		context.Status(http.StatusBadRequest).JSON(
			&fiber.Map{"message": "could not create book"})
		return err
	}

	context.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "book has been added"})
	return nil
}
		
		// func main() {
		// 	log.Print("server has started")
		
		
		
			// db.ConnectDb()
			// pgdb, err := db.ConnectDb()
			// if err != nil {
			// 		log.Printf("error starting the database %v", err)
			// }
			//get the router of the API by passing the db
			// router := api.StartAPI(db.ConnectDb())
			//get the port from the environment variable
			// port := os.Getenv("PORT")
			// app := fiber.New()
			// log.Printf("error starting the database %v", app)
		
			// app.Get("/", func(c *fiber.Ctx) error {
			// 		return c.SendString("Obi app")
			// })
		
			// app.Listen(":3000")
			//pass the router and start listening with the server
			// err = http.ListenAndServe(fmt.Sprintf(":%s", port), router)
			// if err != nil {
			// 		log.Printf("error from router %v\n", err)
			// }
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

// func (r *Repository) SetupRoutes(app *fiber.App) {
	// func SetupRoutes(app *fiber.App) {
	// 	api := app.Group("/api")
	// 	api.Get("/", func(c *fiber.Ctx) error {
	// 		return c.SendString("Hello, Obi111611!")
	// })
	// api.Post("/", r.CreateBook)
	// api.Post("/create_books", r.CreateBook)
	// api.Delete("delete_book/:id", r.DeleteBook)
	// api.Get("/get_books/:id", r.GetBookByID)
	// api.Get("/books", r.GetBooks)
// }

func panicOnError(err error) {
  if err != nil {
    panic(err)
  }
}

func main() {

	fmt.Println("this has started")

	database.ConnectDb()
	// panicOnError(err)
  // defer db.Close()


	// loginUser(c)
	// Signup(c)
	
	app := fiber.New()
	// app.Use(middleware.Logger())
	setupRoutes(app)
	app.Listen(":3000")
}

// func Signup(db *gorm.DB) ( string) { 
  // res, err := signup.Signup(database, &signup.Request{
  //   Email: c.Params("email"),
  //   Password: c.Params("password"),
  // })
	// if err != nil {
  //   switch err.(type){
  //   case *user.EmailDuplicateError:
	// 		msg := fmt.Sprintf("Bad Request: , %s ", err.Error())
  //     return c.SendString(msg)
  //   default:
	// 		msg := fmt.Sprintf("Internal Server Error: , %s ", err.Error())
  //     return c.SendString(msg)
  //   }
  // }
	// msg := fmt.Sprintf("Created: , %s", res.Id)
  // return c.SendString(msg)
// 	return "dfg"
// }


