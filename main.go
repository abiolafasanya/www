package main

import (
    "net/http"

    "github.com/gin-gonic/gin"
)

type book struct {
    ID     string  `json:"id"`
    Title  string  `json:"title"`
    Author string  `json:"author"`
    Price  float64 `json:"price"`
}

var books = []book{
    {ID: "1", Title: "Blue Train", Author: "John Coltrane", Price: 56.99},
    {ID: "2", Title: "Jeru", Author: "Gerry Mulligan", Price: 17.99},
    {ID: "3", Title: "Sarah Vaughan and Clifford Brown", Author: "Sarah Vaughan", Price: 39.99},
}

// getBooks responds with the list of all albums as JSON.
func getBooks(c *gin.Context) {
    c.IndentedJSON(http.StatusOK, books)
}

// addBooks adds an album from JSON received in the request body.
func addBooks(c *gin.Context) {
    var newBook book

    // Call BindJSON to bind the received JSON to
    // newBook.
    if err := c.BindJSON(&newBook); err != nil {
        return
    }

    // Add the new book to the slice.
    books = append(books, newBook)
    c.IndentedJSON(http.StatusCreated, newBook)
}

// getBookByID locates the book whose ID value matches the id
// parameter sent by the client, then returns that album as a response.
func getBookByID(c *gin.Context) {
    id := c.Param("id")

    // Loop through the list of books, looking for
    // an book whose ID value matches the parameter.
    for _, a := range books {
        if a.ID == id {
            c.IndentedJSON(http.StatusOK, a)
            return
        }
    }
    c.IndentedJSON(http.StatusNotFound, gin.H{"message": "book not found"})
}

func main() {
    router := gin.Default()
    router.GET("/books", getBooks)
    router.GET("/books/:id", getBookByID)
    router.POST("/books", addBooks)

    router.Run("localhost:8080")
}