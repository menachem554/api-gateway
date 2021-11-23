package bookstore

import (
	"context"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	pb "github.com/menachem554/Bookstore/proto"
	entity "github.com/menachem554/api-gateway/bookstore/entity.go"

)

// CreateBook : To create the new book
func CreateBook(c *gin.Context) {
	// Get the request as json
	book := PostBook{}
	err := c.ShouldBindJSON(&book)

	if err != nil {
		c.JSON(http.StatusBadRequest, err)
	}

	// Convert to grpc
	book1 := &entity.PostBook{}

	// Send the request to grpc server
	res, err := C.PostBook(context.Background(), &pb.BookRequest{Book: book1})
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, res)
}

// ReadBook : To read the book of the given  iD
func GetBook(c *gin.Context) {
	bookId := c.Param("id")
	req := &pb.GetBookReq{Id: bookId}

	res, err := C.GetBook(c, req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, res)
}

// UpdateBook : To Update the book of the given ID
func UpdateBook(c *gin.Context) {
	bookId := c.Param("id")
	book := &entity.PostBook{}

	err := c.ShouldBindJSON(&book)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
	}

	book1 := &pb.Book{
		BookID:   bookId,
		BookName: book.BookName,
		Category:    book.Category,
		Author:   book.Author,
	}

	res, err := C.UpdateBook(context.Background(), &pb.BookRequest{Book: book1})
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	result := res.GetBook()

	c.JSON(http.StatusOK, result)
}

// DeleteBook: To Delete the book of the given ID
func DeleteBook(c *gin.Context) {
	bookId := c.Param("id")

	req := &pb.GetBookReq{Id: bookId}

	res, err := C.DeleteBook(c, req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"Success": fmt.Sprint(res.Deleted),
	})
}