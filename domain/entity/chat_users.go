package entity

import "time"

type ChatUser struct {
	Id        string    `db:"id" json:"id"`
	Uuid      string    `db:"uuid" json:"uuid"`
	GroupId   uint      `db:"group_id" json:"group_id"`
	UserId    uint      `db:"user_id" json:"user_id"`
	CreatedAt time.Time `db:"created_at" json:"created_at"`
	UpdatedAt time.Time `db:"updated_at" json:"updated_at"`
}

func NewChatUser(
	userId uint,
) *ChatUser {
	return &ChatUser{
		UserId: userId,
	}
}
