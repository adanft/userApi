package router

import (
	"afranco.com/oauth/infrastructure/controller"
	"github.com/gorilla/mux"
)

type UserRouter struct {
	Prefix         string
	userController controller.UserController
}

func NewUserRouter() *UserRouter {
	return &UserRouter{Prefix: "/users", userController: *controller.NewUserController()}
}

func (userRouter *UserRouter) GenerateRoutes(router *mux.Router) {

	subRouter := router.PathPrefix(userRouter.Prefix).Subrouter()

	subRouter.HandleFunc("", userRouter.userController.CreateUser).Methods("POST")
	subRouter.HandleFunc("/{id}", userRouter.userController.GetUser).Methods("GET")
	subRouter.HandleFunc("/{id}", userRouter.userController.DeleteUser).Methods("DELETE")
	subRouter.HandleFunc("", userRouter.userController.GetUsers).Methods("GET")
	subRouter.HandleFunc("/{id}", userRouter.userController.UpdateUser).Methods("PUT")
}
