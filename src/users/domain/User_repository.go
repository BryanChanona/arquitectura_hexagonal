package domain

type Iuser interface {
	SaveUser(user User)(err error)
}