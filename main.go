package main

import (
	"go-soccer/core/service"
	"go-soccer/handler"
	"go-soccer/repository"
	"go-soccer/router"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	host := os.Getenv("APP_HOST")
	port := os.Getenv("APP_PORT")

	teamRepository := repository.NewTeamRepository()
	teamService := service.NewTeamService(teamRepository)
	teamHandler := handler.NewTeamHandler(teamService)

	playerRepository := repository.NewPlayerRepository()
	playerService := service.NewPlayerService(playerRepository, teamRepository)
	playerHandler := handler.NewPlayerHandler(playerService)

	router := router.NewRouter(teamHandler, playerHandler)

	srv := &http.Server{
		Handler:      router,
		Addr:         host + ":" + port,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Println("Soccer API is running!")
	log.Fatal(srv.ListenAndServe())
}
