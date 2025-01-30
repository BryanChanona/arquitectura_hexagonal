package infraestructure

import "database/sql"

//Representa una conexión a la base de datos.
type MySQL struct {
	db *sql.DB
}
//Usamos esta función para crear una instancia de la estructura MySQL
func NewMySQL ( db *sql.DB) *MySQL{
	return &MySQL{db:db}
}


