package repository

import (
	"afranco.com/oauth/domain/entity"
	apierror "afranco.com/oauth/domain/error"
	"afranco.com/oauth/infrastructure/persistence"
	"gorm.io/gorm"
)

type UserRepository struct {
	connection gorm.DB
}

func NewUserRepository() *UserRepository {
	return &UserRepository{connection: *persistence.GetConnection().Connection}
}

func (userRepository *UserRepository) GetUserById(id string) (*entity.User, error) {
	var user entity.User

	res := userRepository.connection.First(&user, id)

	if res.Error != nil {
		return nil, apierror.NewNotFoundError("User not found.")
	}

	return &user, nil
}

func (userRepository *UserRepository) CreateUser(user *entity.User) error {
	res := userRepository.connection.Create(&user)

	if res.Error != nil {
		return apierror.NewConflictError("The username or email already exist.")
	}

	return nil
}

func (userRepository *UserRepository) UpdateUser(user *entity.User, id string) (*entity.User, error) {
	var userDB entity.User

	resFind := userRepository.connection.First(&userDB, id)

	if resFind.Error != nil {
		return nil, resFind.Error
	}

	res := userRepository.connection.Model(&userDB).Updates(user)

	if res.Error != nil {
		return nil, res.Error
	}

	return &userDB, nil
}

func (userRepository *UserRepository) DeleteUser(id string) (bool, error) {
	res := userRepository.connection.Unscoped().Delete(&entity.User{}, id)

	if res.Error != nil {
		return false, res.Error
	}

	if res.RowsAffected == 0 {
		return false, nil
	}

	return true, nil
}

func (userRepository *UserRepository) GetUsers() (*[]entity.User, error) {
	var users []entity.User

	res := userRepository.connection.Find(&users)

	if res.Error != nil {
		return nil, res.Error
	}

	return &users, nil
}
