package controller

import (
	"net/http"

	"github.com/BryanChanona/arquitectura_hexagonal.git/src/users/application"
	"github.com/gin-gonic/gin"
)

type ViewUsersController struct {
	viewUsers *application.ViewUsers
}

func NewViewUsersController(useCase *application.ViewUsers) *ViewUsersController{
	return &ViewUsersController{viewUsers: useCase}
}


func (controller *ViewUsersController) Execute(ctx *gin.Context){
	users, err := controller.viewUsers.Execute()

	if err != nil{
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"users": users})

}
