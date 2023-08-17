package main

import (
	"github.com/gin-gonic/gin"
	"strconv"
)

type Book struct {
	ID     string `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
}

var books []Book

func getBooks(c *gin.Context) {
	c.JSON(200, books)
}

func getBook(c *gin.Context) {
	title := c.Param("title")
	book := getBookByTitle(title)
	if book == nil {
		c.JSON(404, gin.H{"error": "Book not found"})
	} else {
		c.JSON(200, book)
	}
}

func createBook(c *gin.Context) {
	var book Book
	err := c.BindJSON(&book)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	// Generate a unique ID for the book
	book.ID = generateID()

	books = append(books, book)
	c.JSON(201, book)
}

func updateBook(c *gin.Context) {
	title := c.Param("title")
	var updatedBook Book
	err := c.BindJSON(&updatedBook)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	// Find the book by title and update its fields
	for i, book := range books {
		if book.Title == title {
			updatedBook.ID = book.ID
			books[i] = updatedBook
			c.JSON(200, updatedBook)
			return
		}
	}

	c.JSON(404, gin.H{"error": "Book not found"})
}

func deleteBook(c *gin.Context) {
	title := c.Param("title")

	// Find the index of the book by title
	for i, book := range books {
		if book.Title == title {
			// Remove the book from the slice
			books = append(books[:i], books[i+1:]...)
			c.JSON(200, gin.H{"message": "Book deleted"})
			return
		}
	}

	c.JSON(404, gin.H{"error": "Book not found"})
}

func getBookByTitle(title string) *Book {
	for _, book := range books {
		if book.Title == title {
			return &book
		}
	}
	return nil
}

func generateID() string {
	// Implement your own logic to generate a unique ID for the book
	// For simplicity, we'll use a simple counter here
	return strconv.Itoa(len(books) + 1)
}

func main() {
	// Enable CORS
	router := gin.Default()
	router.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(200)
			return
		}
		c.Next()
	})

	// Routes
	router.GET("/books", getBooks)
	router.GET("/books/:title", getBook)
	router.POST("/books", createBook)
	router.PUT("/books/:title", updateBook)
	router.DELETE("/books/:title", deleteBook)

	router.Run(":8080")
}