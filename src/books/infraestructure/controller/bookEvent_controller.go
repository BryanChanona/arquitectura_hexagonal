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
    // Configuramos el timeout de 30 segundos para cada solicitud
    timeout := time.After(30 * time.Second)
    ticker := time.Tick(1 * time.Second) // Revisa cada segundo
    for {
        select {
        case <-timeout:
            // Si no hubo cambios en 30 segundos, cerramos la conexión
            ctx.JSON(http.StatusNoContent, gin.H{"message": "No hubo cambios en el tiempo de espera"})
            return
        case <-ticker:
            // Realizamos la consulta a la base de datos
            result, err := controller.repository.GetAll()
            if err != nil {
                ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error al obtener los libros"})
                return
            }

            currentSize := len(result)

            controller.mu.Lock()
            // Si el tamaño de los libros cambia, respondemos con los nuevos datos
            if currentSize != controller.lastBookCount {
                controller.lastBookCount = currentSize
                controller.mu.Unlock()
                // Enviar la respuesta y reiniciar el ciclo de long polling
                ctx.JSON(http.StatusOK, gin.H{"message": "Hay cambios", "total": currentSize})
                return 
            }
            controller.mu.Unlock()
        }
    }
}

