package entity

import (
	"time"

	"github.com/google/uuid"
)

type Org struct {
	Id        uint      `db:"id" json:"id"`
	Uuid      uuid.UUID `db:"uuid" json:"uuid"`
	Name      string    `db:"name" json:"name"`
	Email      string    `db:"mail" json:"mail"`
	Password  string    `db:"password" json:"password"`
	OrgType   uint      `db:"org_type" json:"org_type"` // 0: guest(default), 1: normal, 2: admin
	CreatedAt time.Time `db:"created_at" json:"created_at"`
	UpdatedAt time.Time `db:"updated_at" json:"updated_at"`

	// relative to other tables
	Users []User `db:"users" json:"users"`

	// not exist in db
	Rank uint `db:"rank" json:"rank"`
}

func NewOrg(
	name string,
	mail string,
	password string,
	orgType uint,
) *Org {
	return &Org{
		Name:     name,
		Email:     mail,
		Password: password,
		OrgType:  orgType,
	}
}
