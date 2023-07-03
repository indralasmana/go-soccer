package handler

import (
	"encoding/json"
	"go-soccer/core/domain"
	"go-soccer/core/port"
	"io"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type (
	PlayerHandler struct {
		playerService port.PlayerService
	}
)

func NewPlayerHandler(playerService port.PlayerService) *PlayerHandler {
	return &PlayerHandler{
		playerService: playerService,
	}
}

func (h *PlayerHandler) GetAllPlayers(w http.ResponseWriter, r *http.Request) {
	players, err := h.playerService.GetAllPlayers()
	if err != nil {
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(map[string]string{"message": err.Error()})
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(players)
}

func (h *PlayerHandler) GetPlayerByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])

	if err != nil {
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"message": "invalid player id"})
		return
	}

	player, err := h.playerService.GetPlayerByID(id)
	if err != nil {
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(map[string]string{"message": err.Error()})
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(player)
}

func (h *PlayerHandler) AddPlayer(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	body, err := io.ReadAll(r.Body)

	if err != nil {
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err.Error())
		return
	}

	var player *domain.Player
	json.Unmarshal(body, &player)

	msg, err := h.playerService.AddPlayer(player)
	if err != nil {
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(err.Error())
		return
	}
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(msg)
}
