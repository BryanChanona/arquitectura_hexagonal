package controller

import (
	"net/http"
	"sync"
	"github.com/BryanChanona/arquitectura_hexagonal.git/src/books/infraestructure"
	"github.com/gin-gonic/gin"
)

type BookPollingController struct {
	repository    *infraestructure.MySQL
	lastBookCount int
	mu            sync.Mutex
}

func NewBookPollingController(repo *infraestructure.MySQL) *BookPollingController{
	return &BookPollingController{
		repository:repo,
		lastBookCount: 0,
	}
}
func (controller *BookPollingController) ShortPollingExecute(ctx *gin.Context){
	result, err := controller.repository.GetAll()
	if err != nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error obteniendo libros"})
        return
    }
	currentUserCount:= len(result)
	
	controller.mu.Lock()
	defer controller.mu.Unlock()

	if currentUserCount != controller.lastBookCount{
		controller.lastBookCount = currentUserCount
	
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