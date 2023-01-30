package routes

import (
	"fmt"
	"net/http"

	"github.com/hidenari-yuda/go-grpc-clean/domain/entity"
	"github.com/hidenari-yuda/go-grpc-clean/infra/database"
	"github.com/hidenari-yuda/go-grpc-clean/infra/di"
	"github.com/hidenari-yuda/go-grpc-clean/usecase"
	"github.com/labstack/echo/v4"
)

type UserRoutes struct {
	Name string
}

func NewUserRoutes() UserRouteFunc {
	return &UserRoutes{
		Name: "UserRoutes",
	}
}

func (r *UserRoutes) SignUp(db *database.DB, firebase usecase.Firebase) func(c echo.Context) error {
	return func(c echo.Context) error {
		var (
			param entity.SignUpParam
		)

		err := bindAndValidate(c, &param)
		if err != nil {
			return c.JSON(http.StatusBadRequest, err)
		}

		tx, _ := db.Begin()
		h := di.InitializeUserHandler(tx, firebase)
		presenter, err := h.SignUp(&param)
		if err != nil {
			tx.Rollback()
			return c.JSON(http.StatusInternalServerError, err)
		}
		tx.Commit()
		c.JSON(http.StatusOK, presenter)
		return nil
	}
}

func (r *UserRoutes) SignIn(db *database.DB, firebase usecase.Firebase) func(c echo.Context) error {
	return func(c echo.Context) error {
		var (
			param entity.SignInParam
		)

		err := bindAndValidate(c, &param)
		if err != nil {
			return c.JSON(http.StatusBadRequest, err)
		}

		h := di.InitializeUserHandler(db, firebase)
		presenter, err := h.SignIn(&param)
		if err != nil {
			err = fmt.Errorf("サインインエラー: %s:%w", err.Error(), entity.ErrRequestError)
			fmt.Println(err)
			return c.JSON(http.StatusInternalServerError, err)
		}

		renderJSON(c, presenter)
		return nil
	}
}
