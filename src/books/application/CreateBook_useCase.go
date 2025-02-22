package application

import "github.com/BryanChanona/arquitectura_hexagonal.git/src/books/domain"

type CreateBook struct {
	db domain.Ibook
}

func NewCreateBook(db domain.Ibook) *CreateBook {
	return &CreateBook{db: db}
}

func (useCase *CreateBook) Execute(book domain.Book) (err error) {
	return useCase.db.SaveBook(book)
}
