package service

import (
	"go-soccer/core/domain"
	"go-soccer/core/port"
)

type (
	PlayerService struct {
		playerRepository port.PlayerRepository
		teamRepository   port.TeamRepository
	}
)

func NewPlayerService(playerRepository port.PlayerRepository, teamRepository port.TeamRepository) port.PlayerService {
	return &PlayerService{
		playerRepository: playerRepository,
		teamRepository:   teamRepository,
	}
}

func (s *PlayerService) GetAllPlayers() ([]*domain.Player, error) {
	return s.playerRepository.GetAllPlayers()
}

func (s *PlayerService) GetPlayerByID(id int) (*domain.Player, error) {
	return s.playerRepository.GetPlayerByID(id)
}

func (s *PlayerService) AddPlayer(player *domain.Player) (*domain.Player, error) {
	_, err := s.teamRepository.GetTeamByID(player.TeamID)
	if err != nil {
		return nil, err
	}
	return s.playerRepository.AddPlayer(player)
}
