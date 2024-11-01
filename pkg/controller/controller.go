package controller

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/CrudOperationUsingGin/pkg/models"

	"github.com/gin-gonic/gin"
)

func CreateBook(c *gin.Context) {
	var newBook models.Book
	if err := c.ShouldBindJSON(&newBook); err != nil {
		log.Fatal(err.Error() + "during create request")
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	b := newBook.CreateBook()
	c.JSON(http.StatusCreated, b)
}

func GetBook(c *gin.Context) {
	newBooks := models.GetBook()

	c.JSON(http.StatusOK, newBooks)
}

func GetBookById(c *gin.Context) {
	bookid := c.Param("bookid")
	fmt.Println(bookid)
	id, err := strconv.ParseInt(bookid, 0, 0)
	if err != nil {
		log.Fatal("error while parsing")
	}
	bookDetails, _ := models.GetBookById(id)
	c.JSON(http.StatusOK, bookDetails)

}

func UpdateBookById(c *gin.Context) {
	var updateBook = &models.Book{}
	bookid := c.Param("bookid")
	if err := c.ShouldBindJSON(&updateBook); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	id, err := strconv.ParseInt(bookid, 0, 0)
	if err != nil {

		log.Fatal("error while parsing")
	}
	bookDetails, db := models.GetBookById(id)
	if updateBook.Name != "" {
		bookDetails.Name = updateBook.Name

	}
	if updateBook.Author != "" {
		bookDetails.Author = updateBook.Author

	}
	if updateBook.Publication != "" {
		bookDetails.Publication = updateBook.Publication

	}
	fmt.Println(bookDetails)
	db.Save(bookDetails)
	c.JSON(http.StatusOK, bookDetails)

}

func DeleteBookById(c *gin.Context) {
	bookid := c.Param("bookid")
	id, err := strconv.ParseInt(bookid, 0, 0)
	if err != nil {
		log.Fatal("error while parsing")
	}
	book := models.DeleteBookById(id)
	c.JSON(http.StatusOK, book)

}
