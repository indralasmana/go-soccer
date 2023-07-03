package repository

import (
	"go-soccer/core/domain"

	"github.com/stretchr/testify/mock"
)

type (
	PlayerRepositoryMock struct {
		Mock mock.Mock
	}
)

func (repository *PlayerRepositoryMock) GetAllPlayers() ([]*domain.Player, error) {
	arguments := repository.Mock.Called()

	var r0 []*domain.Player
	if arguments.Get(0) == nil {
		r0 = nil
	} else {
		r0 = arguments.Get(0).([]*domain.Player)
	}

	var r1 error
	if arguments.Get(1) == nil {
		r1 = nil
	} else {
		r1 = arguments.Error(1)
	}

	return r0, r1
}

func (repository *PlayerRepositoryMock) GetPlayerByID(id int) (*domain.Player, error) {
	arguments := repository.Mock.Called(id)

	var r0 *domain.Player
	if arguments.Get(0) == nil {
		r0 = nil
	} else {
		r0 = arguments.Get(0).(*domain.Player)
	}

	var r1 error
	if arguments.Get(1) == nil {
		r1 = nil
	} else {
		r1 = arguments.Error(1)
	}

	return r0, r1
}

func (repository *PlayerRepositoryMock) AddPlayer(player *domain.Player) (*domain.Player, error) {
	arguments := repository.Mock.Called(player)

	var r0 *domain.Player
	if arguments.Get(0) == nil {
		r0 = nil
	} else {
		r0 = arguments.Get(0).(*domain.Player)
	}

	var r1 error
	if arguments.Get(1) == nil {
		r1 = nil
	} else {
		r1 = arguments.Error(1)
	}

	return r0, r1
}
