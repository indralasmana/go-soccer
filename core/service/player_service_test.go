package service

import (
	"errors"
	"go-soccer/core/domain"
	"go-soccer/repository"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var playerRepository = &repository.PlayerRepositoryMock{Mock: mock.Mock{}}
var playerService = NewPlayerService(playerRepository, teamRepository)

func TestGetAllPlayers(t *testing.T) {
	type (
		param struct {
		}
		expected struct {
			err error
		}
	)

	tests := []struct {
		name     string
		param    param
		expected expected
	}{
		{
			name:     "GetAllPlayers success",
			param:    param{},
			expected: expected{nil},
		},
	}

	playerRepository.Mock.On("GetAllPlayers").Return([]*domain.Player{
		{
			ID:       1,
			TeamID:   1,
			Name:     "Nor Halid",
			Number:   1,
			Position: "GK",
			Nation:   "IDN",
		},
	}, nil)

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			_, err := playerService.GetAllPlayers()
			assert.Equal(t, test.expected.err, err)
		})
	}
}

func TestGetPlayerByID(t *testing.T) {
	type (
		param struct {
			id int
		}
		expected struct {
			err error
		}
	)

	tests := []struct {
		name     string
		param    param
		expected expected
	}{
		{
			name:     "GetPlayerByID success",
			param:    param{id: 1},
			expected: expected{nil},
		},
		{
			name:     "GetPlayerByID failed",
			param:    param{id: 2},
			expected: expected{err: errors.New("player id not found")},
		},
	}

	playerRepository.Mock.On("GetPlayerByID", 1).Return(&domain.Player{
		ID:       1,
		TeamID:   1,
		Name:     "Nor Halid",
		Number:   1,
		Position: "GK",
		Nation:   "IDN",
	}, nil)
	playerRepository.Mock.On("GetPlayerByID", 2).Return(nil, errors.New("player id not found"))

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			_, err := playerService.GetPlayerByID(test.param.id)
			assert.Equal(t, test.expected.err, err)
		})
	}
}

func TestAddPlayer(t *testing.T) {
	type (
		param struct {
			player *domain.Player
		}
		expected struct {
			err error
		}
	)

	tests := []struct {
		name     string
		param    param
		expected expected
	}{
		{
			name: "AddPlayer success",
			param: param{
				player: &domain.Player{
					ID:       1,
					TeamID:   1,
					Name:     "Nor Halid",
					Number:   1,
					Position: "GK",
					Nation:   "IDN",
				},
			},
			expected: expected{nil},
		},
		{
			name: "AddPlayer failed",
			param: param{
				player: &domain.Player{
					ID:       1,
					TeamID:   2,
					Name:     "Nor Halid",
					Number:   1,
					Position: "GK",
					Nation:   "IDN",
				},
			},
			expected: expected{err: errors.New("team id not found")},
		},
	}

	teamRepository.Mock.On("GetTeamByID", 1).Return(&domain.Team{
		ID:    1,
		Name:  "Arema FC",
		Coach: "Joko Susilo",
	}, nil)
	teamRepository.Mock.On("GetTeamByID", 2).Return(nil, errors.New("team id not found"))

	playerRepository.Mock.On("AddPlayer", mock.MatchedBy(func(player *domain.Player) bool { return player.ID == 1 })).Return(&domain.Player{
		ID:       1,
		TeamID:   1,
		Name:     "Nor Halid",
		Number:   1,
		Position: "GK",
		Nation:   "IDN",
	}, nil)

	playerRepository.Mock.On("AddPlayer", mock.MatchedBy(func(player *domain.Player) bool { return player.ID == 1 })).Return(&domain.Player{
		ID:       1,
		TeamID:   1,
		Name:     "Nor Halid",
		Number:   1,
		Position: "GK",
		Nation:   "IDN",
	}, nil)

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			_, err := playerService.AddPlayer(test.param.player)
			assert.Equal(t, test.expected.err, err)
		})
	}
}
