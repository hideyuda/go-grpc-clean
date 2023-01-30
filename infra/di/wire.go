//go:build wireinject
// +build wireinject

package di

import (
	"github.com/google/wire"
	"github.com/hidenari-yuda/go-grpc-clean/handler"
	"github.com/hidenari-yuda/go-grpc-clean/repository"
	"github.com/hidenari-yuda/go-grpc-clean/usecase"
	"github.com/hidenari-yuda/go-grpc-clean/usecase/interactor"
)

var wireSet = wire.NewSet(
	handler.WireSet,
	interactor.WireSet,
	repository.WireSet,
)

/**
	Handler
**/

// User
//
func InitializeUserHandler(db repository.SQLExecuter, fb usecase.Firebase) (h handler.UserHandler) {
	wire.Build(wireSet)
	return
}

/**
	Interactor
**/

// User
//
func InitializeUserInteractor(db repository.SQLExecuter, fb usecase.Firebase) (i interactor.UserInteractor) {
	wire.Build(wireSet)
	return
}
