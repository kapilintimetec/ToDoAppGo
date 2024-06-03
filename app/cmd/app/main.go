package main

import (
	"ToDoAppGo/app/api"
	"ToDoAppGo/app/config"
	"github.com/labstack/echo/v4"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

func main() {
	e := echo.New()

	log.Print("Initializing routes")
	api.InitRoutes(e)

	log.Print("Reading config")
	config.InitConfig()

	log.Print("Connecting to database")

	dsn := "host=db user=user password=password dbname=todoapp port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	_, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Panic(err)
	}

	log.Print("Starting server")

	if err := e.Start(":8080"); err != nil {
		e.Logger.Fatal(err)
	}
}
