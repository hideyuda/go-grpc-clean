package usecase

import "github.com/hidenari-yuda/go-grpc-clean/pb"

type UserRepository interface {
	// Gest API
	// Create
	Create(param *pb.User) error

	// Update
	Update(param *pb.User) error

	// Delete
	Delete(id uint) error

	// Get
	GetById(id uint) (*pb.User, error)

	// auth
	SignIn(mail, password string) (user *pb.User, err error)
	GetByFirebaseId(firebaseId string) (*pb.User, error)

	// admin
	GetAll() ([]*pb.User, error)
}

type ChatGroupRepository interface {
	// Gest API
	// Create
	Create(param *pb.ChatGroup) error

	// Update
	Update(param *pb.ChatGroup) error

	// Delete
	Delete(id uint) error

	// Get
	GetById(id uint) (*pb.ChatGroup, error)

	// get list by user id
	GetListByUserId(userId uint) ([]*pb.ChatGroup, error)
}

type ChatUserRepository interface {
	// Gest API
	// Create
	Create(param *pb.ChatUser) error

	// Delete
	Delete(id uint) error

	// Get
	GetById(id uint) (*pb.ChatUser, error)
	GetListByGroupId(groupId uint) ([]*pb.ChatUser, error)
}

type ChatRepository interface {
	// Gest API
	// Create
	Create(param *pb.Chat) error

	// Delete
	Delete(id uint) error

	// Get
	GetById(id uint) (*pb.Chat, error)
	GetListByGroupId(groupId uint) ([]*pb.Chat, error)
}
