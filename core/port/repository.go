package port

import "go-soccer/core/domain"

type (
	TeamRepository interface {
		GetAllTeams() ([]*domain.Team, error)
		GetTeamByID(id int) (*domain.Team, error)
		AddTeam(team *domain.Team) (*domain.Team, error)
	}

	PlayerRepository interface {
		GetAllPlayers() ([]*domain.Player, error)
		GetPlayerByID(id int) (*domain.Player, error)
		AddPlayer(player *domain.Player) (*domain.Player, error)
	}
)
