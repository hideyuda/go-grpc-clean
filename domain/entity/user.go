package entity

import "github.com/google/uuid"

type User struct {
	Id         uint      `db:"id" json:"id"`
	Uuid       uuid.UUID `db:"uuid" json:"uuid"`
	FirebaseId string    `db:"firebase_id" json:"firebase_id"`
	Name       string    `db:"name" json:"name"`
	Email      string    `db:"email" json:"email"`
	Password   string    `db:"password" json:"password"`
	CreatedAt  string    `db:"created_at" json:"created_at"`
	UpdatedAt  string    `db:"updated_at" json:"updated_at"`
}

func NewUser() *User {
	return &User{}
}

type SignUpParam struct {
	Email    string `db:"email" json:"email" validate:"required"`
	Password string `db:"password" json:"password" validate:"required"`
}

type SignInParam struct {
	Token string `json:"token" validate:"required"`
}
