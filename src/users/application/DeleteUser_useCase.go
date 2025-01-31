package application

import "github.com/BryanChanona/arquitectura_hexagonal.git/src/users/domain"

type DeleteUser struct {
	db domain.Iuser
}

func NewDeleteUser(db domain.Iuser) *DeleteUser{
	return &DeleteUser{db:db}
}

func (useCase *DeleteUser) Execute(id int)(error){
	return useCase.db.DeleteUser(id)
}