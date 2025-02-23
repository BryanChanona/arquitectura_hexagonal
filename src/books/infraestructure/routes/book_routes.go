package routes

import (
	"github.com/BryanChanona/arquitectura_hexagonal.git/src/books/infraestructure/dependencies"
	"github.com/gin-gonic/gin"
)

func BookRouter(router *gin.Engine) {
	routes := router.Group("/books")
	createBookController := dependencies.GetCreateBookController().Execute
	viewBooksController := dependencies.GetViewBooksController().Execute
	deleteBookController := dependencies.GetDeleteBookController().Execute
	updateBookController := dependencies.GetUpdateBookController().Execute
	viewBookbyIdController := dependencies.GetViewBookByIdController().Execute
	shortPolling := dependencies.GetPollingBookController().ShortPollingExecute
	


	routes.POST("/", createBookController)
	routes.GET("/",viewBooksController)
	routes.DELETE("/:id",deleteBookController)
	routes.PUT("/:id",updateBookController)
	routes.GET("/:id",viewBookbyIdController)
	routes.GET("/shortPolling",shortPolling)

}