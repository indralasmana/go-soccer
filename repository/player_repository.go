package repository

import (
	"errors"
	"go-soccer/core/domain"
	"go-soccer/core/port"
	"go-soccer/resource"
)

type (
	PlayerRepository struct {
	}
)

func NewPlayerRepository() port.PlayerRepository {
	return &PlayerRepository{}
}

func (r *PlayerRepository) GetAllPlayers() ([]*domain.Player, error) {
	return resource.Players, nil
}

func (r *PlayerRepository) GetPlayerByID(id int) (*domain.Player, error) {
	for _, player := range resource.Players {
		if player.ID == id {
			return player, nil
		}
	}
	return nil, errors.New("player id not found")
}

func (r *PlayerRepository) AddPlayer(player *domain.Player) (*domain.Player, error) {
	player.ID = len(resource.Players) + 1
	resource.Players = append(resource.Players, player)
	return player, nil
}
