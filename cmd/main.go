package main

import (
	"log"
	"net/http"

	"afranco.com/oauth/domain/entity"
	"afranco.com/oauth/infrastructure/persistence"
	"afranco.com/oauth/infrastructure/router"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {
	// Load environment variables
	err := godotenv.Load("./config/.env")

	if err != nil {
		log.Println(err.Error())
		return
	}

	persistence.GetConnection().Connection.AutoMigrate(entity.User{})
	muxRouter := mux.NewRouter()

	router.NewUserRouter().GenerateRoutes(muxRouter)

	http.ListenAndServe(":3000", muxRouter)
}
