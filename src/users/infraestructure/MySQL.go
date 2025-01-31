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


func (mysql *MySQL) GetAll()([]domain.User,error){
	
	data, err := mysql.db.Query("SELECT * FROM users")

	if err != nil {
		return nil, err
	}
	defer data.Close()

	var users []domain.User
	// Itera sobre todas las filas devueltas por la consulta
	for data.Next(){
		var user domain.User
		err := data.Scan(&user.ID,user.GetName(),user.Email)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	if err := data.Err(); err != nil {
		return nil, err
	}
	return users, nil

}



