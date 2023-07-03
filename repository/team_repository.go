package repository

import (
	"errors"
	"go-soccer/core/domain"
	"go-soccer/core/port"
	"go-soccer/resource"
)

type (
	TeamRepository struct{}
)

func NewTeamRepository() port.TeamRepository {
	return &TeamRepository{}
}

func (r *TeamRepository) GetAllTeams() ([]*domain.Team, error) {
	for _, team := range resource.Teams {
		players := []*domain.Player{}
		for _, player := range resource.Players {
			if player.TeamID == team.ID {
				players = append(players, player)
			}
		}
		team.Players = players
	}

	return resource.Teams, nil
}

func (r *TeamRepository) GetTeamByID(id int) (*domain.Team, error) {
	for _, team := range resource.Teams {
		if team.ID == id {
			players := []*domain.Player{}
			for _, player := range resource.Players {
				if player.TeamID == id {
					players = append(players, player)
				}
			}
			team.Players = players
			return team, nil
		}
	}
	return nil, errors.New("team id not found")
}

func (r *TeamRepository) AddTeam(team *domain.Team) (*domain.Team, error) {
	team.ID = len(resource.Teams) + 1
	resource.Teams = append(resource.Teams, team)
	return team, nil
}
