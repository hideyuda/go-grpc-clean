package entity

import "time"

type Chat struct {
	Id        uint      `db:"id" json:"id"`
	Uuid      string    `db:"uuid" json:"uuid"`
	GroupId   uint      `db:"group_id" json:"group_id"`
	From      uint      `db:"from" json:"from"`
	Content   string    `db:"content" json:"content"`
	CreatedAt time.Time `db:"created_at" json:"created_at"`
	UpdatedAt time.Time `db:"updated_at" json:"updated_at"`
}

func NewChat(
	from uint,
	content string,
	createdAt time.Time,
) *Chat {
	return &Chat{
		From:      from,
		Content:   content,
		CreatedAt: createdAt,
	}
}
