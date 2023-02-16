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
	SignIn(mail, password string) (user *entity.User, err error)
	GetByFirebaseId(firebaseId string) (*entity.User, error)

	// admin
	GetAll() ([]*entity.User, error)
}

type ChatGroupRepository interface {
	// Gest API
	// Create
	Create(param *entity.ChatGroup) error

	// Update
	Update(param *entity.ChatGroup) error

	// Delete
	Delete(id uint) error

	// Get
	GetById(id uint) (*entity.ChatGroup, error)
}

type ChatUserRepository interface {
	// Gest API
	// Create
	Create(param *entity.ChatUser) error

	// Delete
	Delete(id uint) error

	// Get
	GetById(id uint) (*entity.ChatUser, error)
	GetListByGroupId(groupId uint) ([]*entity.ChatUser, error)
}

type ChatRepository interface {
	// Gest API
	// Create
	Create(param *entity.Chat) error

	// Delete
	Delete(id uint) error

	// Get
	GetById(id uint) (*entity.Chat, error)
	GetListByGroupId(groupId uint) ([]*entity.Chat, error)
}
