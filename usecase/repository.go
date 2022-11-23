package usecase

import "github.com/hidenari-yuda/umerun-resume/domain/entity"

type UserRepository interface {
	// Gest API
	SignUp(param *entity.SignUpParam) error
	GetByFirebaseId(firebaseId string) (*entity.User, error)
}
