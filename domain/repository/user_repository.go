package irepository

import (
	"afranco.com/oauth/domain/entity"
)

type IUserRepository interface {
	GetUserById(id string) (*entity.User, error)
	CreateUser(user *entity.User) error
	UpdateUser(user *entity.User, id string) (*entity.User, error)
	DeleteUser(id string) (bool, error)
	GetUsers() (*[]entity.User, error)
}
