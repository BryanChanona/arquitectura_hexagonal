package controller

import (
	"net/http"
	"sync"

	"github.com/BryanChanona/arquitectura_hexagonal.git/src/users/infraestructure"
	"github.com/gin-gonic/gin"
)

type UserPollingController struct {
	repository *infraestructure.MySQL
	lastUserCount int
	mu sync.Mutex
}

func NewUserPollingController(repo *infraestructure.MySQL) *UserPollingController{
	return &UserPollingController{
		repository:repo,
		lastUserCount: 0,
	}
}

func (controller *UserPollingController) Execute(ctx *gin.Context){
	result, err := controller.repository.GetAll()
	if err != nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error obteniendo tickets"})
        return
    }
	currentUserCount:= len(result)
	
	controller.mu.Lock()
	defer controller.mu.Unlock()

	if currentUserCount != controller.lastUserCount{
		controller.lastUserCount = currentUserCount
	
		ctx.JSON(http.StatusOK, gin.H{
            "message": "Hay cambios",
            "data":   currentUserCount,
        })


	}else {
        ctx.JSON(http.StatusOK, gin.H{
            "message": "No hay cambios",
        })
	}

}