package main

import (
	dependenciesUser "github.com/BryanChanona/arquitectura_hexagonal.git/src/users/infraestructure/dependencies"
	userRoutesResource "github.com/BryanChanona/arquitectura_hexagonal.git/src/users/infraestructure/routes"
	"github.com/gin-gonic/gin"
	dependenciesBook "github.com/BryanChanona/arquitectura_hexagonal.git/src/books/infraestructure/dependencies"
	bookRoutesResource "github.com/BryanChanona/arquitectura_hexagonal.git/src/books/infraestructure/routes"
)

func main() {
	r := gin.Default()
	//User
		dependenciesUser.Init()
		userRoutesResource.UserRouter(r)
	
	//Book
		dependenciesBook.Init()
		bookRoutesResource.BookRouter(r)
		

		r.Run()
}