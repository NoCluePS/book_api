package main

import (
	"fiber_intro/book"
	"fiber_intro/database"
	"fiber_intro/models"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)


func setupRoutes(app *fiber.App) {
	app.Get("/api/book", book.GetBooks)
	app.Get("/api/book/:id", book.GetBook)
	app.Post("/api/book", book.NewBook)
	app.Delete("/api/book/:id", book.DeleteBook)
	app.Patch("/api/book/:id", book.UpdateBook)
	app.Delete("/api/book", book.ClearDB)
}

func initDatabase() {
	var err error
	database.DBConn, err = gorm.Open(sqlite.Open("books.db"), &gorm.Config{})
	
	if err != nil {
		panic("failed to connect database")
	}
	fmt.Println("Connection Opened to Database")

	database.DBConn.AutoMigrate(&models.Book{})
	fmt.Println("Successfully migrated modal structs")
}

func main() {
	app := fiber.New();
	initDatabase()

	setupRoutes(app)
	app.Listen(":3000")
}