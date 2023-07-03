package resource

import (
	"go-soccer/core/domain"
)

var Teams = []*domain.Team{
	{
		ID:      1,
		Name:    "Barito Putera",
		Coach:   "Rahmad Darmawan",
		Players: []*domain.Player{},
	},
	{
		ID:      2,
		Name:    "Persija Jakarta",
		Coach:   "Thomas Doll",
		Players: []*domain.Player{},
	},
}
