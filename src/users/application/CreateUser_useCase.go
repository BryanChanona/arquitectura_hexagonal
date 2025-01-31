package application

import "github.com/BryanChanona/arquitectura_hexagonal.git/src/users/domain"

type CreateUser struct {
	db domain.Iuser
}

func NewCreateUser(db domain.Iuser) *CreateUser {
	return &CreateUser{db: db}

}

func (useCase *CreateUser) Execute(user domain.User) (err error) {
	return useCase.db.SaveUser(user)
}
