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
	Email      string    `db:"email" json:"email"`
	Password   string    `db:"password" json:"password"`
	UserType   uint      `db:"user_type" json:"user_type"` // 0: guest(default), 1: normal, 2: admin
	CreatedAt  time.Time `db:"created_at" json:"created_at"`
	UpdatedAt  time.Time `db:"updated_at" json:"updated_at"`

	// relative to other tables
	// Org_Id     uint        `db:"org_id" json:"org_id"`
	ChatGroups []ChatGroup `db:"chat_groups" json:"chat_groups"`

	Medias   []Media   `db:"medias" json:"medias"`
	Articles []Article `db:"articles" json:"articles"`
	Keywords []Keyword `db:"keywords" json:"keywords"`

	// not exist in db
	// Rank uint `db:"rank" json:"rank"`
}

func NewUser(
	name string,
	email string,
	password string,
	userType uint,
) *User {
	return &User{
		Name:     name,
		Email:    email,
		Password: password,
		UserType: userType,
	}
}

type SignUpParam struct {
	Email    string `db:"email" json:"email" validate:"required"`
	Password string `db:"password" json:"password" validate:"required"`
}

type SignInParam struct {
	Token string `json:"token" validate:"required"`
}

type Media struct {
	Id        uint   `db:"id" json:"id"`
	Uuid      string `db:"uuid" json:"uuid"`
	UserId    uint   `db:"user_id" json:"user_id"`
	Title     string `db:"title" json:"title"`
	Url       string `db:"url" json:"url"`
	Type      uint   `db:"type" json:"type"`
	CreatedAt string `db:"created_at" json:"created_at"`
	UpdatedAt string `db:"updated_at" json:"updated_at"`
	IsDeleted bool   `db:"is_deleted" json:"is_deleted"`
}

type MediaCompetitor struct {
	Id        uint   `db:"id" json:"id"`
	Uuid      string `db:"uuid" json:"uuid"`
	MediaId   uint   `db:"media_id" json:"media_id"`
	Title     string `db:"title" json:"title"`
	Url       string `db:"url" json:"url"`
	Type      uint   `db:"type" json:"type"`
	CreatedAt string `db:"created_at" json:"created_at"`
	UpdatedAt string `db:"updated_at" json:"updated_at"`
	IsDeleted bool   `db:"is_deleted" json:"is_deleted"`
}

type Article struct {
	Id        uint   `db:"id" json:"id"`
	Uuid      string `db:"uuid" json:"uuid"`
	UserId    uint   `db:"user_id" json:"user_id"`
	Title     string `db:"title" json:"title"`
	MetaTitle string `db:"meta_title" json:"meta_title"`
	MetaDisc  string `db:"meta_disc" json:"meta_disc"`
	Url       string `db:"url" json:"url"`
	CreatedAt string `db:"created_at" json:"created_at"`
	UpdatedAt string `db:"updated_at" json:"updated_at"`
	IsDeleted bool   `db:"is_deleted" json:"is_deleted"`

	Paragraphs []Paragraph `db:"paragraphs" json:"paragraphs"`
}

type Paragraph struct {
	Id        uint   `db:"id" json:"id"`
	Uuid      string `db:"uuid" json:"uuid"`
	ArticleId uint   `db:"article_id" json:"article_id"`
	Headline  string `db:"headline" json:"headline"`
	Content   string `db:"content" json:"content"`
	CreatedAt string `db:"created_at" json:"created_at"`
	UpdatedAt string `db:"updated_at" json:"updated_at"`
	IsDeleted bool   `db:"is_deleted" json:"is_deleted"`
}

type Keyword struct {
	Id            uint   `db:"id" json:"id"`
	Uuid          string `db:"uuid" json:"uuid"`
	UserId        uint   `db:"user_id" json:"user_id"`
	Keyword       string `db:"keyword" json:"keyword"`
	Bolume        uint   `db:"bolume" json:"bolume"`
	Rank          uint   `db:"rank" json:"rank"`
	SeoDifficulty uint   `db:"seo_difficulty" json:"seo_difficulty"`
	Url           string `db:"url" json:"url"`
	CreatedAt     string `db:"created_at" json:"created_at"`
	UpdatedAt     string `db:"updated_at" json:"updated_at"`
	IsDeleted     bool   `db:"is_deleted" json:"is_deleted"`

	// relative to other tables
	KeywordTrends []KeywordTrend `db:"keyword_trends" json:"keyword_trends"`
}

type KeywordTrend struct {
	Id        uint   `db:"id" json:"id"`
	Uuid      string `db:"uuid" json:"uuid"`
	KeywordId uint   `db:"keyword_id" json:"keyword_id"`
	Keyword   string `db:"keyword" json:"keyword"`
	Bolume    uint   `db:"bolume" json:"bolume"`
	Rank      uint   `db:"rank" json:"rank"`
	CreatedAt string `db:"created_at" json:"created_at"`
	UpdatedAt string `db:"updated_at" json:"updated_at"`
	IsDeleted bool   `db:"is_deleted" json:"is_deleted"`
}

type KeywordCompetitor struct {
	Id        uint   `db:"id" json:"id"`
	Uuid      string `db:"uuid" json:"uuid"`
	KeywordId uint   `db:"keyword_id" json:"keyword_id"`
	Url       string `db:"url" json:"url"`
	Rank      uint   `db:"rank" json:"rank"`
	CreatedAt string `db:"created_at" json:"created_at"`
	UpdatedAt string `db:"updated_at" json:"updated_at"`
	IsDeleted bool   `db:"is_deleted" json:"is_deleted"`
}
