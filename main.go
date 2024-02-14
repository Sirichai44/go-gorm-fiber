package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

const (
	host     = "localhost"  // or the Docker service name if running in another container
	port     = 5432         // default PostgreSQL port
	user     = "myuser"     // as defined in docker-compose.yml
	password = "mypassword" // as defined in docker-compose.yml
	dbname   = "mydatabase" // as defined in docker-compose.yml
)

func main() {
	dsn := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             time.Second, // Slow SQL threshold
			LogLevel:                  logger.Info, // Log level
			IgnoreRecordNotFoundError: true,        // Ignore ErrRecordNotFound error for logger        // Don't include params in the SQL log
			Colorful:                  true,        // Disable color
		},
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		panic("failed to connect database")
	}
	fmt.Println("Connection Opened to Database")
	
	db.AutoMigrate(&Book{})
	fmt.Println("Database Migrated")

	//setup fiber
	app := fiber.New()

	app.Get("/books", func(c *fiber.Ctx) error {
		books,err := getBooks(db)
		if err != nil {
			return c.SendStatus(fiber.StatusBadRequest)
		}
		return c.JSON(books)
	})

	app.Get("/books/:id", func(c *fiber.Ctx) error {
		id,err := strconv.Atoi(c.Params("id"))
		if err != nil {
			return c.SendStatus(fiber.StatusBadRequest)
		}
		book := getBook(db, id)
		return c.JSON(book)
	})

	app.Post("/books", func(c *fiber.Ctx) error {
		book := new(Book)
		if err := c.BodyParser(book); err != nil {
			return c.SendStatus(fiber.StatusBadRequest)
		}
		err := createBook(db, book)
		if err != nil {
			return c.SendStatus(fiber.StatusBadRequest)
		}
		return c.JSON((fiber.Map{"message": "Book Created Successfully"}))
	})

	app.Put("/books/:id", func(c *fiber.Ctx) error {
		id,err := strconv.Atoi(c.Params("id"))
		if err != nil {
			return c.SendStatus(fiber.StatusBadRequest)
		}
		book := new(Book)
		if err := c.BodyParser(book); err != nil {
			return c.SendStatus(fiber.StatusBadRequest)
		}
		book.ID = uint(id)
		err = updateBook(db, book)
		if err != nil {
			return c.SendStatus(fiber.StatusBadRequest)
		}
		return c.JSON((fiber.Map{"message": "Book Updated Successfully"}))
	})

	app.Delete("/books/:id", func(c *fiber.Ctx) error {
		id,err := strconv.Atoi(c.Params("id"))
		if err != nil {
			return c.SendStatus(fiber.StatusBadRequest)
		}
		deleteBook(db, id)
		return c.JSON((fiber.Map{"message": "Book Deleted Successfully"}))
	})

	app.Listen(":8080")

}
