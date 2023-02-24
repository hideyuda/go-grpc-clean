package handler

import (
	"fmt"
	"strconv"

	"github.com/hidenari-yuda/go-grpc-clean/domain/config"
	"github.com/hidenari-yuda/go-grpc-clean/domain/entity"
	"github.com/labstack/echo/v4"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/status"
)

func GetFirebaseToken(c echo.Context) string {
	token := c.Request().Header.Get("FirebaseAuthorization")
	return token
}

// func bindAndValidate(c echo.Context, obj interface{}) (err error) {
// 	err = c.Bind(obj)
// 	if err != nil {
// 		return
// 	}
// 	return validator.New().Struct(obj)
// }

// func renderJSON(c echo.Context, p presenter.Presenter) {
// 	c.JSON(p.StatusCode(), p.Data())
// }

// // ファイル送信用
// func renderFile(c echo.Context, filePath string) error {
// 	err := c.File(filePath)
// 	if err != nil {
// 		return err
// 	}
// 	return nil
// }

func handleError(err error) error {
	var (
		code, message = entity.ErrorInfo(err)
		statusCode    = strconv.Itoa(code)
		// path          = ctx.Path()
		// method        = ctx.Request().Method
		errText = err.Error()
	)

	text := fmt.Sprintf("*ERROR* env:%s status: %s message: %s error_text: %s", config.App.Env, statusCode, message, errText)

	stat, _ := status.FromError(err)
	stat, _ = stat.WithDetails(&errdetails.DebugInfo{
		Detail:       text,
		StackEntries: []string{},
	})

	err = stat.Err()

	return err
}
