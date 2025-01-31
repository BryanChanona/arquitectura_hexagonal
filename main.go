package main

import (
	"log"
	"github.com/BryanChanona/arquitectura_hexagonal.git/src/helpers"
)

func main() {

	db, err := helpers.ConnectDB()

	if err != nil {
		log.Fatalf("Error al conectar a la base de datos: %v", err)
		

	}

	db.Close()

}