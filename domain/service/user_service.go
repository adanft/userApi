package iservice

import "afranco.com/oauth/domain/entity"

type IUserService interface {
	GetUser(id string) (*entity.User, error)
	CreateUser(user *entity.User) error
	UpdateUser(user *entity.User, id string) (*entity.User, error)
	DeleteUser(id string) (bool, error)
	GetUsers() (*[]entity.User, error)
}
