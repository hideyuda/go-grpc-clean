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

type ChatRepository interface {
	// Gest API
	// Create
	Create(param *pb.Chat) error

	// Update
	Update(param *pb.Chat) error

	// Delete
	Delete(id uint) error

	// Get
	GetById(id uint) (*pb.Chat, error)

	// get list by user id
	GetListByGroupId(groupId uint) ([]*pb.Chat, error)
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

type MediaRepository interface {
	// Gest API
	// Create
	Create(param *pb.Media) error

	// Update
	Update(param *pb.Media) error

	// Delete
	Delete(id int64) error

	// Get
	GetById(id int64) (*pb.Media, error)

	// get list by user id
	GetListByUserId(userId int64) ([]*pb.Media, error)

	// get list by type
	GetListByType(mediaType int64) ([]*pb.Media, error)

	// get all
	GetAll() ([]*pb.Media, error)
}

type KeywordRepository interface {
	// Gest API
	// Create
	Create(param *pb.Keyword) error

	// Update
	Update(param *pb.Keyword) error

	// Delete
	Delete(id int64) error

	// Get
	GetById(id int64) (*pb.Keyword, error)

	// get list by user id
	GetListByMediaId(mediaId int64) ([]*pb.Keyword, error)
}

type ScriptRepository interface {
	// Gest API
	// Create
	Create(param *pb.Script) error

	// Update
	Update(param *pb.Script) error

	// Delete
	Delete(id int64) error

	// Get
	GetById(id int64) (*pb.Script, error)

	// get list by media id
	GetListByMediaId(mediaId int64) ([]*pb.Script, error)

	// get all
	GetAll() ([]*pb.Script, error)
}