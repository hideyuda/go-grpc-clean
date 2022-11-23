package interactor

import (
	"fmt"

	"github.com/hidenari-yuda/go-docker-template/domain/entity"
	"github.com/hidenari-yuda/go-docker-template/usecase"
)

type UserInteractor interface {
	// Gest API
	SignUp(input SignUpInput) (output SignUpOutput, err error)
	SignIn(input SignInInput) (output SignInOutput, err error)
}

type UserInteractorImpl struct {
	firebase       usecase.Firebase
	userRepository usecase.UserRepository
}

func NewUserInteractorImpl(
	fb usecase.Firebase,
	uR usecase.UserRepository,
) UserInteractor {
	return &UserInteractorImpl{
		firebase:       fb,
		userRepository: uR,
	}
}

type SignUpInput struct {
	Param *entity.SignUpParam
}

type SignUpOutput struct {
	Ok bool
}

func (i *UserInteractorImpl) SignUp(input SignUpInput) (output SignUpOutput, err error) {
	// ユーザー登録
	err = i.userRepository.SignUp(input.Param)
	if err != nil {
		return output, err
	}

	output.Ok = true

	return output, nil
}

type SignInInput struct {
	Param *entity.SignInParam
}

type SignInOutput struct {
	User *entity.User
}

func (i *UserInteractorImpl) SignIn(input SignInInput) (output SignInOutput, err error) {

	firebaseId, err := i.firebase.VerifyIDToken(input.Param.Token)
	if err != nil {
		return output, err
	}

	fmt.Println("exported firebaseToken is:", input.Param.Token)
	fmt.Println("exported firebaseId is:", firebaseId)

	// ユーザー登録
	output.User, err = i.userRepository.GetByFirebaseId(firebaseId)
	if err != nil {
		return output, err
	}

	return output, nil

}

type GetByFirebaseTokenInput struct {
	Token string
}

type GetByFirebaseTokenOutput struct {
	User *entity.User
}
