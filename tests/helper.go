package tests

import (
	"github.com/hidenari-yuda/go-grpc-clean/infra/database"
)

type Helper struct {
	dbm *database.Db
}

func NewHelper() *Helper {
	dbm := database.NewDb()
	err := dbm.MigrateUp("../../.migrations")
	if err != nil {
		panic(err)
	}

	return &Helper{
		dbm: dbm,
	}
}

func (h *Helper) ClearAllTables() {
	h.dbm.Exec(
		`DELETE FROM users`,
		`DROP TABLE users`,
	)
}
