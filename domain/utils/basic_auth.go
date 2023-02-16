package utils

import (
	"crypto/hmac"
	"crypto/sha256"
	"crypto/subtle"
	"encoding/hex"
	"fmt"

	"github.com/hidenari-yuda/go-grpc-clean/domain/config"
	"github.com/labstack/echo/v4"
)

// type BasicAuth struct {
// 	cfg config.Config
// }

// func NewBasicAuth(cfg config.Config) *BasicAuth {
// 	return &BasicAuth{
// 		cfg: cfg,
// 	}
// }

func BasicAuthValidator(username, password string, c echo.Context) (bool, error) {
	var (
		u, p string
	)

	for i, name := range config.App.BasicUsers {
		if username == name {
			u = config.App.BasicUsers[i]
			p = config.App.BasicPasswords[i]
		}
	}
	if p == "" || u == "" {
		return false, nil
	}

	mac := hmac.New(sha256.New, []byte(config.App.BasicSecret))
	mac.Write([]byte(p))
	expectedPass := hex.EncodeToString(mac.Sum(nil))

	fmt.Println("username: ", username)
	fmt.Println("password:", password)
	fmt.Println("expectedPass:", expectedPass)

	if subtle.ConstantTimeCompare([]byte(username), []byte(u)) == 1 &&
		subtle.ConstantTimeCompare([]byte(password), []byte(expectedPass)) == 1 {
		return true, nil
	}

	return false, nil
}
