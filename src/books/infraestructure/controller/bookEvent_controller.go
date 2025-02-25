package controller

import (
	"net/http"
	"sync"
	"time"

	"github.com/BryanChanona/arquitectura_hexagonal.git/src/books/infraestructure"
	"github.com/gin-gonic/gin"
)

type BookPollingController struct {
	repository    *infraestructure.MySQL
	lastBookCount int
	mu            sync.Mutex
}

func NewBookPollingController(repo *infraestructure.MySQL) *BookPollingController {
	return &BookPollingController{
		repository:    repo,
		lastBookCount: 0,
	}
}
func (controller *BookPollingController) ShortPollingExecute(ctx *gin.Context) {
	result, err := controller.repository.GetAll()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error obteniendo libros"})
		return
	}
	currentUserCount := len(result)

	controller.mu.Lock()
	defer controller.mu.Unlock()

	if currentUserCount != controller.lastBookCount {
		controller.lastBookCount = currentUserCount

		ctx.JSON(http.StatusOK, gin.H{
			"message": "Hay cambios",
			"data":    currentUserCount,
		})

	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "No hay cambios",
		})
	}

}




func (controller *BookPollingController) LongPolling(ctx *gin.Context) {
	timeout := time.After(15 * time.Second)
	ticker := time.NewTicker(1 * time.Second) // Revisa cada segundo
	defer ticker.Stop()

	for {
		select {
		case <-timeout:
			ctx.JSON(http.StatusOK, gin.H{"message": "Sin cambios", "total": controller.lastBookCount})
			return

		case <-ticker.C:
			result, err := controller.repository.GetAll()
			if err != nil {
				ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error al obtener los libros"})
				return
			}

			currentSize := len(result)

			controller.mu.Lock()
			if currentSize != controller.lastBookCount {
				controller.lastBookCount = currentSize
				controller.mu.Unlock()
				ctx.JSON(http.StatusOK, gin.H{"message": "Hay cambios", "total": currentSize})
				return
			}
			controller.mu.Unlock()
			 
		}
	}
}