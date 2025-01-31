package infraestructure

import (
	"database/sql"
	"fmt"

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
		err := data.Scan(&user.ID,&user.Name,&user.Email)
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

func (mysql *MySQL) DeleteUser(id int)(error){

	query := "DELETE FROM users WHERE id = ?"
	result, err := mysql.db.Exec(query, id)
	if err != nil {
		return err
	}
	// Verificar si realmente se eliminó algún registro
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return fmt.Errorf("no se encontró el usuario con ID %d", id)
	}
	fmt.Println("Producto eliminado correctamente")
	return nil

}



