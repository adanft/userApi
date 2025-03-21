package persistence

import (
	"log"
	"os"

	"sync"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type OnceConnection struct {
	Connection *gorm.DB
}

var onceConnection *OnceConnection
var once sync.Once

func GetConnection() *OnceConnection {
	if onceConnection == nil {
		once.Do(
			func() {
				gormConn, err := gorm.Open(mysql.Open(os.Getenv("DATABASE_URL")), &gorm.Config{})
				if err != nil {
					log.Panic("DB Connection error: " + err.Error())
				} else {
					log.Println("DB Connection successfully!!!")
				}

				onceConnection = &OnceConnection{Connection: gormConn}
			})
	} else {
		log.Println("DB Connection already exist.")
	}

	return onceConnection
}
