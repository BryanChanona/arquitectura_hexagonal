package routes

import (
	"github.com/BryanChanona/arquitectura_hexagonal.git/src/books/infraestructure/dependencies"
	"github.com/gin-gonic/gin"
)

func BookRouter(router *gin.Engine) {
	routes := router.Group("/books")
	createBookController := dependencies.GetCreateBookController().Execute
	viewBooksController := dependencies.GetViewBooksController().Execute


	routes.POST("/", createBookController)
	routes.GET("/",viewBooksController)
	
}