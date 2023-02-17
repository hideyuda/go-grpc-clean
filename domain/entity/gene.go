package entity

// import (
// 	"time"

// 	"github.com/google/uuid"
// )

// type User struct {
// 	Id         uint      `db:"id" json:"id"`
// 	Uuid       uuid.UUID `db:"uuid" json:"uuid"`
// 	FirebaseId string    `db:"firebase_id" json:"firebase_id"`
// 	Name       string    `db:"name" json:"name"`
// 	Email      string    `db:"email" json:"email"`
// 	Password   string    `db:"password" json:"password"`
// 	UserType   uint      `db:"user_type" json:"user_type"` // 0: guest(default), 1: normal, 2: admin
// 	CreatedAt  time.Time `db:"created_at" json:"created_at"`
// 	UpdatedAt  time.Time `db:"updated_at" json:"updated_at"`
// 	IsDeleted  bool      `db:"is_deleted" json:"is_deleted"`

// 	// relative to other tables
// 	// Org_Id     uint        `db:"org_id" json:"org_id"`
// 	ChatGroups []ChatGroup `db:"chat_groups" json:"chat_groups"`

// 	Generateds []Generated `db:"generated" json:"generated"`

// 	// not exist in db
// 	// Rank uint `db:"rank" json:"rank"`
// }

// func NewUser(
// 	name string,
// 	email string,
// 	password string,
// 	userType uint,
// ) *User {
// 	return &User{
// 		Name:     name,
// 		Email:    email,
// 		Password: password,
// 		UserType: userType,
// 	}
// }

// type SignUpParam struct {
// 	Email    string `db:"email" json:"email" validate:"required"`
// 	Password string `db:"password" json:"password" validate:"required"`
// }

// type SignInParam struct {
// 	Token string `json:"token" validate:"required"`
// }

// type Generated struct {
// 	Id        uint      `db:"id" json:"id"`
// 	Uuid      uuid.UUID `db:"uuid" json:"uuid"`
// 	UserId    uint      `db:"user_id" json:"user_id"`
// 	Type      uint      `db:"type" json:"type"` // 0: text, 1: image
// 	Content   string    `db:"content" json:"content"`
// 	CreatedAt time.Time `db:"created_at" json:"created_at"`
// 	UpdatedAt time.Time `db:"updated_at" json:"updated_at"`
// 	IsDeleted bool      `db:"is_deleted" json:"is_deleted"`

// 	Words []Word `db:"words" json:"words"`
// }

// type Word struct {
// 	Id          uint      `db:"id" json:"id"`
// 	Uuid        uuid.UUID `db:"uuid" json:"uuid"`
// 	GeneratedId uint      `db:"generated_id" json:"generated_id"`
// 	Word        string    `db:"word" json:"word"`
// 	CreatedAt   time.Time `db:"created_at" json:"created_at"`
// 	UpdatedAt   time.Time `db:"updated_at" json:"updated_at"`
// 	IsDeleted   bool      `db:"is_deleted" json:"is_deleted"`
// }
