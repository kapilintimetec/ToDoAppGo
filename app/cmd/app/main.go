package main

import (
	"ToDoAppGo/app/api"
	"ToDoAppGo/app/config"
	"ToDoAppGo/app/internal/db"
	"github.com/labstack/echo/v4"
	"log"
)

func main() {
	log.Print("Loading config")
	cfg, err := config.GetConfig()
	if err != nil {
		log.Fatalf("Error loading config: %v", err)
	}
	log.Print("Config loaded")
	e := echo.New()

	log.Print("Connecting to database")
	_, err = db.GetInstance(cfg)
	if err != nil {
		log.Fatal(err)

	}
	log.Print("Database connected")

	log.Print("Initializing routes")
	api.InitRoutes(e)

	log.Print("Starting server")
	e.Logger.Fatal(e.Start(":8080"))
}
