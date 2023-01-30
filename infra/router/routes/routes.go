package routes

import (
	"github.com/hidenari-yuda/go-grpc-clean/infra/database"
	"github.com/hidenari-yuda/go-grpc-clean/usecase"
	"github.com/labstack/echo/v4"
)

type UserRouteFunc interface {
	SignUp(db *database.DB, firebase usecase.Firebase) func(c echo.Context) error
	SignIn(db *database.DB, firebase usecase.Firebase) func(c echo.Context) error
}
