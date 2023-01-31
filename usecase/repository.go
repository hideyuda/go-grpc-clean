package usecase

import "github.com/hidenari-yuda/go-grpc-clean/domain/entity"

type UserRepository interface {
	// Gest API
	// Create
	Create(param *entity.User) error

	// Update
	Update(param *entity.User) error
	UpdateColumnStr(lineUserId, column, value string) error
	UpdateColumnInt(lineUserId, column string, value int) error

	// Get
	GetById(id uint) (*entity.User, error)
	SignIn(email, password string) (user *entity.User, err error)
	GetByFirebaseId(firebaseId string) (*entity.User, error)

	// admin
	GetAll() ([]*entity.User, error)
}
