package interactor

import (
	"fmt"

	"github.com/hidenari-yuda/go-grpc-clean/pb"
	"github.com/hidenari-yuda/go-grpc-clean/usecase"
)

type UserInteractor interface {
	// Gest API
	// Create
	Create(user *pb.User) (*pb.User, error)

	// Update
	Update(user *pb.User) (bool, error)

	// Delete
	Delete(id uint) (bool, error)

	// Get
	GetById(id uint) (*pb.User, error)

	// auth
	SignIn(param *pb.SignInRequest) (*pb.User, error)

	// admin API
	GetAll() ([]*pb.User, error)
}

type UserInteractorImpl struct {
	firebase usecase.Firebase
	userRepository usecase.UserRepository
}

func NewUserInteractorImpl(
	fb usecase.Firebase,
	uR usecase.UserRepository,
) UserInteractor {
	return &UserInteractorImpl{
		firebase: fb,
		userRepository: uR,
	}
}

// Create
func (i *UserInteractorImpl) Create(user *pb.User) (*pb.User, error) {

	// ユーザー登録
	err := i.userRepository.Create(user)
	if err != nil {
		return user, err
	}

	return user, nil
}

// Update
func (i *UserInteractorImpl) Update(user *pb.User) (bool, error) {

	// ユーザー登録
	err := i.userRepository.Update(user)
	if err != nil {
		return false, err
	}

	return true, nil
}

// delete
func (i *UserInteractorImpl) Delete(id uint) (bool, error) {

	err := i.userRepository.Delete(id)
	if err != nil {
		return false, err
	}

	return true, nil
}

// GetById
func (i *UserInteractorImpl) GetById(id uint) (*pb.User, error) {
	var (
		user *pb.User
		err  error
	)

	// ユーザー登録
	user, err = i.userRepository.GetById(id)
	if err != nil {
		fmt.Println("error is:", err)
		return user, err
	}

	return user, nil
}

// SignIn
func (i *UserInteractorImpl) SignIn(param *pb.SignInRequest) (*pb.User, error) {
	var (
		user *pb.User
		err  error
	)

	firebaseId, err := i.firebase.VerifyIDToken(param.Token)
	if err != nil {
		return user, err
	}

	fmt.Println("exported firebaseToken is:", param.Token)
	fmt.Println("exported firebaseId is:", firebaseId)

	// ユーザー登録
	user, err = i.userRepository.GetByFirebaseId(firebaseId)
	if err != nil {
		return user, err
	}

	return user, nil

}

// GetAll
func (i *UserInteractorImpl) GetAll() ([]*pb.User, error) {
	var (
		users []*pb.User
		err   error
	)

	users, err = i.userRepository.GetAll()
	if err != nil {
		return users, err
	}

	// i.userClient.DetectTextFromImage()

	return users, nil
}
