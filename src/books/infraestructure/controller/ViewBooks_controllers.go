package controller

import (
	"net/http"

	"github.com/BryanChanona/arquitectura_hexagonal.git/src/books/application"
	"github.com/gin-gonic/gin"
)

type ViewBooksController struct {
	viewBooks *application.ViewBooks
}

func NewViewBooksController(useCase *application.ViewBooks) *ViewBooksController {
	return &ViewBooksController{viewBooks: useCase}
}

func (controller *ViewBooksController) Execute(ctx *gin.Context) {
	books, err := controller.viewBooks.Execute()

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"books": books})
}
