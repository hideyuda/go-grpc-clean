package entity

import "time"

type Message struct {
	Id        string    `db:"id" json:"id"`
	Uuid      string    `db:"uuid" json:"uuid"`
	From      string    `db:"from" json:"from"`
	Content   string    `db:"content" json:"content"`
	CreatedAt time.Time `db:"created_at" json:"created_at"`
	UpdatedAt time.Time `db:"updated_at" json:"updated_at"`
}

type Messages []Message

func NewMessage(from, content string, createdAt time.Time) *Message {
	return &Message{
		From:      from,
		Content:   content,
		CreatedAt: createdAt,
	}
}
