package service

import (
	"afranco.com/oauth/domain/entity"
	irepository "afranco.com/oauth/domain/repository"
	"afranco.com/oauth/infrastructure/repository"
)

type UserService struct {
	userRepository irepository.IUserRepository
}

func NewUserService() *UserService {
	return &UserService{userRepository: repository.NewUserRepository()}
}

func (userService *UserService) GetUser(id string) (*entity.User, error) {

	user, err := userService.userRepository.GetUserById(id)

	return user, err
}

func (userService *UserService) CreateUser(user *entity.User) error {
	res := userService.userRepository.CreateUser(user)

	return res
}

func (userService *UserService) UpdateUser(user *entity.User, id string) (*entity.User, error) {
	res, err := userService.userRepository.UpdateUser(user, id)

	return res, err
}

func (userService *UserService) DeleteUser(id string) (bool, error) {
	res, err := userService.userRepository.DeleteUser(id)

	return res, err
}

func (userService *UserService) GetUsers() (*[]entity.User, error) {
	users, res := userService.userRepository.GetUsers()

	return users, res
}
