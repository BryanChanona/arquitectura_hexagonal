package controller

import (
	"net/http"

	"github.com/BryanChanona/arquitectura_hexagonal.git/src/users/application"
	"github.com/BryanChanona/arquitectura_hexagonal.git/src/users/domain"
	"github.com/gin-gonic/gin"
)


type CreateUserController struct {
	createUser *application.CreateUser
}

func NewCreateUserController(useCase *application.CreateUser) *CreateUserController{
	return &CreateUserController{createUser: useCase}

}


func (controller *CreateUserController) CreateProduct(ctx *gin.Context){
	var user domain.User

	if err := ctx.ShouldBindJSON(&user); err != nil{
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err := controller.createUser.Execute(user)

	if err != nil{
		ctx.JSON(500, gin.H{"error": err.Error()})
	}else {
		ctx.JSON(http.StatusCreated, gin.H{"message": "User created"})
	}



}


