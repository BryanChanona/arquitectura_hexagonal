package infraestructure

import (
	"database/sql"

	"github.com/BryanChanona/arquitectura_hexagonal.git/src/users/domain"
)

//Representa una conexión a la base de datos.
type MySQL struct {
	db *sql.DB
}
//Usamos esta función para crear una instancia de la estructura MySQL
func NewMySQL ( db *sql.DB) *MySQL{
	return &MySQL{db:db}
}


func (mysql *MySQL)SaveUser(user domain.User) (err error){
	sentenciaPreparada, err := mysql.db.Prepare("INSERT INTO users ( name, email) VALUES (?,?)")

	if err != nil{
		return err
	}

	defer sentenciaPreparada.Exec()

	_,err = sentenciaPreparada.Exec(user.GetName(),user.GetEmail())

	if err != nil {
		return err
	}
	return nil
	
}


