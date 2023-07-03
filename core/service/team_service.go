package service

import (
	"go-soccer/core/domain"
	"go-soccer/core/port"
)

type (
	TeamService struct {
		teamRepository port.TeamRepository
	}
)

func NewTeamService(teamRepository port.TeamRepository) port.TeamService {
	return &TeamService{
		teamRepository: teamRepository,
	}
}

func (s *TeamService) GetAllTeams() ([]*domain.Team, error) {
	return s.teamRepository.GetAllTeams()
}

func (s *TeamService) GetTeamByID(id int) (*domain.Team, error) {
	return s.teamRepository.GetTeamByID(id)
}

func (s *TeamService) AddTeam(team *domain.Team) (*domain.Team, error) {
	return s.teamRepository.AddTeam(team)
}
