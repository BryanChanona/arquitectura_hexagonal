package controller

import (
	"net/http"

	"github.com/BryanChanona/arquitectura_hexagonal.git/src/books/application"
	"github.com/BryanChanona/arquitectura_hexagonal.git/src/books/domain"
	"github.com/gin-gonic/gin"
)

type CreateBookController struct {
	createBook *application.CreateBook
}

func NewCreateBookController(useCase *application.CreateBook) *CreateBookController {
	return &CreateBookController{createBook: useCase}
}

func (controller *CreateBookController) Execute(ctx *gin.Context) {
	var book domain.Book

	if err := ctx.ShouldBindJSON(&book); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Execute the use case to create the book
	err := controller.createBook.Execute(book)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	} else {
		ctx.JSON(http.StatusCreated, gin.H{"message": "Book created"})
	}
}
