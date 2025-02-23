package routes

import (
	"github.com/BryanChanona/arquitectura_hexagonal.git/src/users/infraestructure/dependencies"
	"github.com/gin-gonic/gin"
)

func UserRouter(router *gin.Engine){
	routes := router.Group("/users")
	createUserController := dependencies.GetCreateUserController().Execute
	viewUsersController := dependencies.GetViewUsersController().Execute
	deleteUserController := dependencies.GetDeleteUserController().Execute
	updateUserController := dependencies.GetUpdateUserController().Execute
	viewUserByIdController := dependencies.GetViewUserByIdController().Execute

	shortPolling := dependencies.GetPollingUserController().Execute
	longPolling := dependencies.GetPollingUserController().LongPolling

	routes.POST("/",createUserController)
	routes.GET("/",viewUsersController)
	routes.DELETE("/:id",deleteUserController)
	routes.PUT("/:id", updateUserController)
	routes.GET("/shortPolling",shortPolling)
	routes.GET("/:id",viewUserByIdController)
	routes.GET("/longPolling",longPolling)
}