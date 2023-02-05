package entity

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	Id         uint      `db:"id" json:"id"`
	Uuid       uuid.UUID `db:"uuid" json:"uuid"`
	FirebaseId string    `db:"firebase_id" json:"firebase_id"`
	Name       string    `db:"name" json:"name"`
	Mail       string    `db:"mail" json:"mail"`
	Password   string    `db:"password" json:"password"`
	UserType   uint      `db:"user_type" json:"user_type"` // 0: guest(default), 1: normal, 2: admin
	CreatedAt  time.Time `db:"created_at" json:"created_at"`
	UpdatedAt  time.Time `db:"updated_at" json:"updated_at"`

	// relative to other tables
	Org_Id     uint        `db:"org_id" json:"org_id"`
	ChatGroups []ChatGroup `db:"chat_groups" json:"chat_groups"`

	// not exist in db
	Rank uint `db:"rank" json:"rank"`
}

func NewUser(
	name string,
	mail string,
	password string,
	userType uint,
) *User {
	return &User{
		Name:     name,
		Mail:     mail,
		Password: password,
		UserType: userType,
	}
}

type SignUpParam struct {
	Mail     string `db:"mail" json:"mail" validate:"required"`
	Password string `db:"password" json:"password" validate:"required"`
}

type SignInParam struct {
	Token string `json:"token" validate:"required"`
}
