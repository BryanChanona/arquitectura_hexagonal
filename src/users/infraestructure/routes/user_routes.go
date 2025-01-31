package routes

import (
	"github.com/BryanChanona/arquitectura_hexagonal.git/src/users/infraestructure/dependencies"
	"github.com/gin-gonic/gin"
)

func UserRouter(router *gin.Engine){
	routes := router.Group("/users")
	createUserController := dependencies.GetCreateUserController().Execute
	viewUsersController := dependencies.GetViewUsersController().Execute

	routes.POST("/",createUserController)
	routes.GET("/",viewUsersController)
}