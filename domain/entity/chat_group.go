package entity

import "time"

type ChatGroup struct {
	Id        string    `db:"id" json:"id"`
	Uuid      string    `db:"uuid" json:"uuid"`
	Name      string    `db:"name" json:"name"`
	Desc      string    `db:"desc" json:"desc"`
	CreatedAt time.Time `db:"created_at" json:"created_at"`
	UpdatedAt time.Time `db:"updated_at" json:"updated_at"`

	Chats     []Chat     `json:"chats"`
	ChatUsers []ChatUser `json:"users"`
}

func NewChatGroup(
	name string,
) *ChatGroup {
	return &ChatGroup{
		Name: name,
	}
}
