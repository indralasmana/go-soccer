package router

import (
	"go-soccer/handler"
	"net/http"

	"github.com/gorilla/mux"
)

func NewRouter(teamHandler *handler.TeamHandler, playerHandler *handler.PlayerHandler) *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/teams", teamHandler.GetAllTeams).Methods(http.MethodGet)
	router.HandleFunc("/teams/{id}", teamHandler.GetTeamByID).Methods(http.MethodGet)
	router.HandleFunc("/teams", teamHandler.AddTeam).Methods(http.MethodPost)
	router.HandleFunc("/players", playerHandler.GetAllPlayers).Methods(http.MethodGet)
	router.HandleFunc("/players/{id}", playerHandler.GetPlayerByID).Methods(http.MethodGet)
	router.HandleFunc("/players", playerHandler.AddPlayer).Methods(http.MethodPost)
	return router
}
