package application

import "github.com/BryanChanona/arquitectura_hexagonal.git/src/books/domain"

type ViewBooks struct {
	db domain.Ibook
}

func NewViewBooks(db domain.Ibook) *ViewBooks {
	return &ViewBooks{db: db}
}

func (useCase *ViewBooks) Execute() ([]domain.Book, error) {
	return useCase.db.GetAll()
}
