package usecase

import "github.com/hidenari-yuda/go-grpc-clean/domain/entity"

type UserRepository interface {
	// Gest API
	SignUp(param *entity.SignUpParam) error
	GetByFirebaseId(firebaseId string) (*entity.User, error)
}
