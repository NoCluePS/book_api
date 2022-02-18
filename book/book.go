package book

import (
	"fiber_intro/database"
	"fiber_intro/models"

	"github.com/gofiber/fiber/v2"
)

func GetBooks(c *fiber.Ctx) error {
	db := database.DBConn
	var books []models.Book
	db.Find(&books)
	c.JSON(books)

	return nil
}

func GetBook(c *fiber.Ctx) error {
	id := c.Params("id");
	db := database.DBConn
	var book models.Book

	db.Find(&book, id)

	if (book.Title == "") {
		c.Status(400).JSON(fiber.Map{
			"error": "Didn't find book with that ID",
		})
	}

	c.JSON(book)

	return nil
}

func NewBook(c *fiber.Ctx) error {
	db := database.DBConn
	
	book := new(models.Book)
	if err := c.BodyParser(book); err != nil {
		c.Status(503).JSON(fiber.Map{
			"error": err,
		})
		return err;
	}

	db.Create(&book)
	c.JSON(book)

	return nil
}

func DeleteBook(c *fiber.Ctx) error {
	id := c.Params("id");
	db := database.DBConn
	var book models.Book
	db.First(&book, id)

	if (book.Title == "") {
		c.Status(404).JSON(fiber.Map{
			"err": "No book found with given ID",
		})
	}
	
	db.Delete(&book)
	c.JSON(fiber.Map{
		"message": "Deleted book",
	})

	return nil
}

func UpdateBook(c *fiber.Ctx) error {
	id := c.Params("id")
	db := database.DBConn
	var book models.Book
	db.First(&book, id)

	if (book.Title == "") {
		c.Status(404).JSON(fiber.Map{
			"error": "Book not found",
		})
	}

	newBook := new(models.Book)
	if err := c.BodyParser(newBook); err != nil {
		c.Status(400).JSON(fiber.Map{
			"error": err,
		})
		return err;
	}

	book.Author = newBook.Author
	book.Title = newBook.Title
	book.Rating = newBook.Rating

	db.Save(&book)
	c.JSON(book)

	return nil
}

func ClearDB(c *fiber.Ctx) error {
	db := database.DBConn
	db.Exec("DELETE FROM books")

	return nil
}