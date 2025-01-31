package dependencies

import (
	"log"

	"github.com/BryanChanona/arquitectura_hexagonal.git/src/helpers"
	"github.com/BryanChanona/arquitectura_hexagonal.git/src/books/application"
	"github.com/BryanChanona/arquitectura_hexagonal.git/src/books/infraestructure"
	"github.com/BryanChanona/arquitectura_hexagonal.git/src/books/infraestructure/controller"
)

var (
	mySQL infraestructure.MySQL
)

func Init(){
	db, err := helpers.ConnectDB()
	if err != nil {
		log.Fatalf("Error al conectar a la base de datos: %v", err)
	}
	mySQL =*infraestructure.NewMySQL(db)

}

func GetCreateBookController() *controller.CreateBookController {
	caseCreateBook := application.NewCreateBook(&mySQL)
	return controller.NewCreateBookController(caseCreateBook)
}
