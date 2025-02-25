package controller

import (
	"net/http"
	"sync"
	"time"

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
        ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error obteniendo usuarios"})
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

func (controller *UserPollingController) LongPolling(ctx *gin.Context) {
    // Configuramos el timeout de 30 segundos para cada solicitud
    timeout := time.After(15 * time.Second)
    ticker := time.NewTicker(1 * time.Second) 
    defer ticker.Stop()
    // Revisa cada segundo
    for {
        select {
        case <-timeout:
            // Si no hubo cambios en 30 segundos, cerramos la conexión
            ctx.JSON(http.StatusOK, gin.H{"message": "No hubo cambios en el tiempo de espera"})
            return
        case <-ticker.C:
            // Realizamos la consulta a la base de datos
            result, err := controller.repository.GetAll()
            if err != nil {
                ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error al obtener los usuarios"})
                return
            }

            currentSize := len(result)

            controller.mu.Lock()
           
            if currentSize != controller.lastUserCount {
                controller.lastUserCount = currentSize
                controller.mu.Unlock()
                // Enviar la respuesta y reiniciar el ciclo de long polling
                ctx.JSON(http.StatusOK, gin.H{"message": "Hay cambios", "total": currentSize})
                return 
            }
            controller.mu.Unlock()
        }
    }
}