package interactor

import (
	"fmt"

	"github.com/hidenari-yuda/go-grpc-clean/domain/entity"
	"github.com/hidenari-yuda/go-grpc-clean/usecase"
)

type UserInteractor interface {
	// Gest API
	// Create
	Create(input CreateInput) (CreateOutput, error)

	// Update
	Update(input UpdateInput) (UpdateOutput, error)

	// Get
	GetById(input GetByIdInput) (GetByIdOutput, error)
	SignIn(input SignInInput) (SignInOutput, error)
	GetByFirebaseToken(input GetByFirebaseTokenInput) (GetByFirebaseTokenOutput, error)

	// admin API
	GetAll() (GetAllOutput, error)
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

// Create
type CreateInput struct {
	Param *entity.User
}

type CreateOutput struct {
	Ok bool
}

func (i *UserInteractorImpl) Create(input CreateInput) (CreateOutput, error) {
	var (
		output CreateOutput
		err    error
	)

	// ユーザー登録
	err = i.userRepository.Create(input.Param)
	if err != nil {
		return output, err
	}

	output.Ok = true

	return output, nil
}

// Update
type UpdateInput struct {
	Param *entity.User
}

type UpdateOutput struct {
	Ok bool
}

func (i *UserInteractorImpl) Update(input UpdateInput) (UpdateOutput, error) {
	var (
		output UpdateOutput
		err    error
	)

	// ユーザー登録
	err = i.userRepository.Create(input.Param)
	if err != nil {
		return output, err
	}

	output.Ok = true

	return output, nil
}

// Get
type GetByIdInput struct {
	Id uint
}

type GetByIdOutput struct {
	User *entity.User
}

func (i *UserInteractorImpl) GetById(input GetByIdInput) (GetByIdOutput, error) {
	var (
		output GetByIdOutput
		err    error
	)

	// ユーザー登録
	output.User, err = i.userRepository.GetById(input.Id)
	if err != nil {
		fmt.Println("error is:", err)
		return output, err
	}

	return output, nil
}

type SignInInput struct {
	Param *entity.SignInParam
}

type SignInOutput struct {
	User *entity.User
}

func (i *UserInteractorImpl) SignIn(input SignInInput) (SignInOutput, error) {
	var (
		output SignInOutput
		err    error
	)

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

func (i *UserInteractorImpl) GetByFirebaseToken(input GetByFirebaseTokenInput) (GetByFirebaseTokenOutput, error) {
	var (
		output GetByFirebaseTokenOutput
		err    error
	)

	firebaseId, err := i.firebase.VerifyIDToken(input.Token)
	if err != nil {
		return output, err
	}

	fmt.Println("exported firebaseId is:", firebaseId)

	user, err := i.userRepository.GetByFirebaseId(firebaseId)
	if err != nil {
		return output, err
	}

	fmt.Println("exported user is:", user)

	output.User = user

	return output, nil
}

type GetAllOutput struct {
	Users []*entity.User
}

func (i *UserInteractorImpl) GetAll() (GetAllOutput, error) {
	var (
		output GetAllOutput
		err    error
	)

	users, err := i.userRepository.GetAll()
	if err != nil {
		return output, err
	}

	output.Users = users

	return output, nil
}
