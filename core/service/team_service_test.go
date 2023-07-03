package service

import (
	"errors"
	"go-soccer/core/domain"
	"go-soccer/repository"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var teamRepository = &repository.TeamRepositoryMock{Mock: mock.Mock{}}
var teamService = NewTeamService(teamRepository)

func TestGetAllTeams(t *testing.T) {
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
			name:     "GetAllTeams success",
			param:    param{},
			expected: expected{nil},
		},
	}

	teamRepository.Mock.On("GetAllTeams").Return([]*domain.Team{
		{
			ID:    1,
			Name:  "Arema FC",
			Coach: "Joko Susilo",
		},
	}, nil)

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			_, err := teamService.GetAllTeams()
			assert.Equal(t, test.expected.err, err)
		})
	}
}

func TestGetTeamByID(t *testing.T) {
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
			name:     "GetTeamByID success",
			param:    param{id: 1},
			expected: expected{nil},
		},
		{
			name:     "GetTeamByID failed",
			param:    param{id: 2},
			expected: expected{err: errors.New("team id not found")},
		},
	}

	teamRepository.Mock.On("GetTeamByID", 1).Return(&domain.Team{
		ID:    1,
		Name:  "Arema FC",
		Coach: "Joko Susilo",
	}, nil)
	teamRepository.Mock.On("GetTeamByID", 2).Return(nil, errors.New("team id not found"))

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			_, err := teamService.GetTeamByID(test.param.id)
			assert.Equal(t, test.expected.err, err)
		})
	}
}

func TestAddTeam(t *testing.T) {
	type (
		param struct {
			team *domain.Team
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
			name: "AddTeam success",
			param: param{
				team: &domain.Team{
					ID:    1,
					Name:  "Arema FC",
					Coach: "Joko Susilo",
				},
			},
			expected: expected{nil},
		},
	}

	teamRepository.Mock.On("AddTeam", mock.MatchedBy(func(team *domain.Team) bool { return team.ID == 1 })).Return(&domain.Team{
		ID:    1,
		Name:  "Arema FC",
		Coach: "Joko Susilo",
	}, nil)

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			_, err := teamService.AddTeam(test.param.team)
			assert.Equal(t, test.expected.err, err)
		})
	}
}
