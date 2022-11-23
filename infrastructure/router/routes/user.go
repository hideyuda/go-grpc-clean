package routes

import (
	"fmt"
	"net/http"

	"github.com/hidenari-yuda/umerun-resume/domain/entity"
	"github.com/hidenari-yuda/umerun-resume/infrastructure/database"
	"github.com/hidenari-yuda/umerun-resume/infrastructure/di"
	"github.com/hidenari-yuda/umerun-resume/usecase"
	"github.com/labstack/echo/v4"
)

type UserRouteFunc interface {
	SignUp(db *database.DB, firebase usecase.Firebase) func(c echo.Context) error
	SignIn(db *database.DB, firebase usecase.Firebase) func(c echo.Context) error
	GetByFirebaseToken(db *database.DB, firebase usecase.Firebase) func(c echo.Context) error
	UploadResume(db *database.DB, firebase usecase.Firebase) func(c echo.Context) error
}

type UserRoutes struct {
	UserRouteFunc
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

func (r *UserRoutes) GetByFirebaseToken(db *database.DB, firebase usecase.Firebase) func(c echo.Context) error {
	return func(c echo.Context) error {
		var (
			firebaseToken = GetFirebaseToken(c)
		)

		h := di.InitializeUserHandler(db, firebase)
		presenter, err := h.GetByFirebaseToken(firebaseToken)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err)
		}

		c.JSON(http.StatusOK, presenter)
		return nil
	}
}

func (r *UserRoutes) DetectTextFromJobSeekerResume(db *database.DB, firebase usecase.Firebase) func(c echo.Context) error {
	return func(c echo.Context) error {

		// fileParam, err := c.FormFile("file")
		// if err != nil {
		// 	wrapped := fmt.Errorf("ファイルの受け取りエラー: %s:%w", err.Error(), entity.ErrRequestError)
		// 	renderJSON(c, presenter.NewErrorJsonPresenter(wrapped))
		// 	return wrapped
		// }

		// file, err := fileParam.Open()
		// if err != nil {
		// 	wrapped := fmt.Errorf("ファイルが開けません: %s:%w", err.Error(), entity.ErrRequestError)
		// 	renderJSON(c, presenter.NewErrorJsonPresenter(wrapped))
		// 	return wrapped
		// }

		// fmt.Println(fileParam.Filename)
		// fmt.Println(fileParam.Size)
		// fmt.Println(fileParam.Header)

		// defer file.Close()

		filePath := "../../../.public/images/townwork_template.png"

		// filePath := fmt.Sprintf("%v-%v.png", time.Now(), fileParam.Filename)

		// ioutil.WriteFile(filePath, []byte(file), 0644)

		h := di.InitializeUserHandler(db, firebase)
		presenter, err := h.DetectTextFromJobSeekerResume(filePath)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err)
		}
		c.JSON(http.StatusOK, presenter)
		return nil
	}
}
