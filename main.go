package main

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

type book struct {
    ID     string `json:"id"`
    Title  string `json:"title"`
    Author string `json:"author"`
    Qty    int    `json:"qty"`
}

var books = []book{
    {ID: "1", Title: "Golang", Author: "Google", Qty: 10},
    {ID: "2", Title: "Java", Author: "Oracle", Qty: 20},
    {ID: "3", Title: "Python", Author: "Python Software Foundation", Qty: 30},
}

func checkoutBook(c *gin.Context) {
    id, ok := c.GetQuery("id")

    if !ok {
        c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "id is required"})
        return
    }

    book, err := getBooksByID(id)

    if err != nil {
        c.IndentedJSON(http.StatusNotFound, gin.H{"message": "book not found"})
        return
    }

    if book.Qty <= 0 {
        c.IndentedJSON(http.StatusNotFound, gin.H{"message": "book out of stock"})
        return
    }

    book.Qty--
    c.IndentedJSON(http.StatusOK, book)
}

func returnBook(c *gin.Context) {
    id, ok := c.GetQuery("id")

    if !ok {
        c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "id is required"})
        return
    }

    book, err := getBooksByID(id)

    if err != nil {
        c.IndentedJSON(http.StatusNotFound, gin.H{"message": "book not found"})
        return
    }

    book.Qty++
    c.IndentedJSON(http.StatusOK, book)
}

func getBooks(c *gin.Context) {
    c.IndentedJSON(http.StatusOK, books)
}

func bookByID(c *gin.Context) {
    id := c.Param("id")
    book, err := getBooksByID(id)

    if err != nil {
        c.IndentedJSON(http.StatusNotFound, gin.H{"message": "book not found"})
        return
    } else {
        c.IndentedJSON(http.StatusOK, book)
        return
    }
}

func getBooksByID(id string) (*book, error) {
    for i, b := range books {
        if b.ID == id {
            return &books[i], nil
        }
    }

    return nil, errors.New("book not found")
}

func createBook(c *gin.Context) {
    var newBook book

    if err := c.BindJSON(&newBook); err != nil {
        c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "invalid request"})
        return
    }

    books = append(books, newBook)
    c.IndentedJSON(http.StatusCreated, newBook)
}

func main() {
    router := gin.Default()
    router.GET("/books", getBooks)
    router.POST("/books", createBook)
    router.GET("/books/:id", bookByID)
    router.PATCH("/checkout", checkoutBook)
    router.PATCH("/return", returnBook)
    router.Run("localhost:8080")
}
