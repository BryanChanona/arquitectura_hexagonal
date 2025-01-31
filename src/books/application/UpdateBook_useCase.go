package application

import "github.com/BryanChanona/arquitectura_hexagonal.git/src/books/domain"

type UpdateBook struct {
	db domain.Ibook
}

func NewUpdateBook(db domain.Ibook) *UpdateBook {
	return &UpdateBook{db: db}
}

func (useCase *UpdateBook) Execute(id int, book domain.Book) (err error) {
	return useCase.db.UpdateBook(id, book)
}
