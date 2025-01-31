package main

import (
	dependenciesUser "github.com/BryanChanona/arquitectura_hexagonal.git/src/users/infraestructure/dependencies"
	"github.com/BryanChanona/arquitectura_hexagonal.git/src/users/infraestructure/routes"
	"github.com/gin-gonic/gin"
)

func main() {
		dependenciesUser.Init()
		r := gin.Default()
		routes.UserRouter(r)
		r.Run()

}