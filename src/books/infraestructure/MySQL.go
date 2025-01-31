package infraestructure

import (
	"database/sql"

	"github.com/BryanChanona/arquitectura_hexagonal.git/src/books/domain"
)

// Representa una conexión a la base de datos.
type MySQL struct {
	db *sql.DB
}

// Usamos esta función para crear una instancia de la estructura MySQL
func NewMySQL(db *sql.DB) *MySQL {
	return &MySQL{db: db}
}

func (mysql *MySQL) SaveBook(book domain.Book) (err error) {
	// Preparar la sentencia SQL para insertar el libro
	sentenciaPreparada, err := mysql.db.Prepare("INSERT INTO books (title, author) VALUES (?, ?)")

	if err != nil {
		return err
	}

	// Asegurarse de que la sentencia preparada se ejecute correctamente al final
	defer sentenciaPreparada.Exec()

	// Ejecutar la sentencia preparada con los valores del libro
	_, err = sentenciaPreparada.Exec(book.GetTitle(), book.GetAuthor())

	if err != nil {
		return err
	}
	return nil
}

