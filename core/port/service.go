package port

import "go-soccer/core/domain"

type (
	TeamService interface {
		GetAllTeams() ([]*domain.Team, error)
		GetTeamByID(id int) (*domain.Team, error)
		AddTeam(team *domain.Team) (*domain.Team, error)
	}

	PlayerService interface {
		GetAllPlayers() ([]*domain.Player, error)
		GetPlayerByID(id int) (*domain.Player, error)
		AddPlayer(player *domain.Player) (*domain.Player, error)
	}
)
