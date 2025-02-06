package main

import (
	dependenciesBook "github.com/BryanChanona/arquitectura_hexagonal.git/src/books/infraestructure/dependencies"
	bookRoutesResource "github.com/BryanChanona/arquitectura_hexagonal.git/src/books/infraestructure/routes"
	"github.com/BryanChanona/arquitectura_hexagonal.git/src/helpers"
	dependenciesUser "github.com/BryanChanona/arquitectura_hexagonal.git/src/users/infraestructure/dependencies"
	userRoutesResource "github.com/BryanChanona/arquitectura_hexagonal.git/src/users/infraestructure/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	dependenciesUser.Init()
	dependenciesBook.Init()


	r := gin.Default()

	helpers.InitCORS(r)
	
	// Configurar rutas
	
	userRoutesResource.UserRouter(r)
	bookRoutesResource.BookRouter(r)

	r.Run()  // Especifica explícitamente el puerto si es necesario
}
