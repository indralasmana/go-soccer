package repository

import (
	"go-soccer/core/domain"

	"github.com/stretchr/testify/mock"
)

type (
	TeamRepositoryMock struct {
		Mock mock.Mock
	}
)

func (repository *TeamRepositoryMock) GetAllTeams() ([]*domain.Team, error) {
	arguments := repository.Mock.Called()

	var r0 []*domain.Team
	if arguments.Get(0) == nil {
		r0 = nil
	} else {
		r0 = arguments.Get(0).([]*domain.Team)
	}

	var r1 error
	if arguments.Get(1) == nil {
		r1 = nil
	} else {
		r1 = arguments.Error(1)
	}

	return r0, r1
}

func (repository *TeamRepositoryMock) GetTeamByID(id int) (*domain.Team, error) {
	arguments := repository.Mock.Called(id)

	var r0 *domain.Team
	if arguments.Get(0) == nil {
		r0 = nil
	} else {
		r0 = arguments.Get(0).(*domain.Team)
	}

	var r1 error
	if arguments.Get(1) == nil {
		r1 = nil
	} else {
		r1 = arguments.Error(1)
	}

	return r0, r1
}

func (repository *TeamRepositoryMock) AddTeam(team *domain.Team) (*domain.Team, error) {
	arguments := repository.Mock.Called(team)

	var r0 *domain.Team
	if arguments.Get(0) == nil {
		r0 = nil
	} else {
		r0 = arguments.Get(0).(*domain.Team)
	}

	var r1 error
	if arguments.Get(1) == nil {
		r1 = nil
	} else {
		r1 = arguments.Error(1)
	}

	return r0, r1
}
