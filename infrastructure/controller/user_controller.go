package controller

import (
	"encoding/json"
	"net/http"

	"afranco.com/oauth/application/dto"
	"afranco.com/oauth/application/service"
	"afranco.com/oauth/domain/entity"
	"afranco.com/oauth/domain/response"
	iservice "afranco.com/oauth/domain/service"
	"github.com/gorilla/mux"
)

type UserController struct {
	userService iservice.IUserService
}

func NewUserController() *UserController {
	return &UserController{userService: service.NewUserService()}
}

func (userController *UserController) CreateUser(res http.ResponseWriter, req *http.Request) {
	var user entity.User
	json.NewDecoder(req.Body).Decode(&user)

	err := userController.userService.CreateUser(&user)

	if err != nil {
		res.WriteHeader(http.StatusConflict)
		res.Write([]byte(err.Error()))
		return
	}

	var userResponse dto.UserResponse = dto.UserResponse{
		Id:        int(user.ID),
		Username:  user.Username,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
		Avatar:    user.Avatar,
	}

	var response = response.NewBaseResponse(userResponse)

	json.NewEncoder(res).Encode(&response)
}

func (userController *UserController) GetUser(res http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)

	user, err := userController.userService.GetUser(params["id"])

	if err != nil {
		res.WriteHeader(http.StatusNotFound)
		res.Write([]byte(err.Error()))
		return
	}

	var userResponse dto.UserResponse = dto.UserResponse{
		Id:        int(user.ID),
		Username:  user.Username,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
		Avatar:    user.Avatar,
	}

	json.NewEncoder(res).Encode(&userResponse)

}

func (userController *UserController) UpdateUser(res http.ResponseWriter, req *http.Request) {
	var user entity.User
	json.NewDecoder(req.Body).Decode(&user)
	params := mux.Vars(req)

	userDB, err := userController.userService.UpdateUser(&user, params["id"])

	if err != nil {
		res.WriteHeader(http.StatusConflict)
		res.Write([]byte(err.Error()))
		return
	}

	var userResponse dto.UserResponse = dto.UserResponse{
		Id:        int(userDB.ID),
		Username:  userDB.Username,
		FirstName: userDB.FirstName,
		LastName:  userDB.LastName,
		Email:     userDB.Email,
		Avatar:    userDB.Avatar,
	}

	var response = response.NewBaseResponse(userResponse)

	json.NewEncoder(res).Encode(&response)
}

func (userController *UserController) DeleteUser(res http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)

	isDeleted, err := userController.userService.DeleteUser(params["id"])

	if err != nil {
		res.Write([]byte(err.Error()))
	}

	var response = response.NewBaseResponse[any](nil)
	response.Success = isDeleted

	response.Message = "The data was deleted successfully!!!"

	json.NewEncoder(res).Encode(&response)
}

func (userController *UserController) GetUsers(res http.ResponseWriter, req *http.Request) {
	users, err := userController.userService.GetUsers()

	if err != nil {
		res.Write([]byte(err.Error()))
	}

	var usersResponse []dto.UserResponse = []dto.UserResponse{}

	for _, user := range *users {
		user := dto.UserResponse{
			Id:        int(user.ID),
			Username:  user.Username,
			FirstName: user.FirstName,
			LastName:  user.LastName,
			Email:     user.Email,
			Avatar:    user.Avatar,
		}
		usersResponse = append(usersResponse, user)
	}

	var response = response.NewBaseResponse(usersResponse)
	response.Message = "The data was requested successfully!!!"

	json.NewEncoder(res).Encode(&response)
}
