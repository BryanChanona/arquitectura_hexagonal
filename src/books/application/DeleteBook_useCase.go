package application

import "github.com/BryanChanona/arquitectura_hexagonal.git/src/books/domain"

type DeleteBook struct {
	db domain.Ibook
}

func NewDeleteBook(db domain.Ibook) *DeleteBook {
	return &DeleteBook{db: db}
}

func (useCase *DeleteBook) Execute(id int) error {
	return useCase.db.DeleteBook(id)
}
