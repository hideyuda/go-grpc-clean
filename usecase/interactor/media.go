package interactor

import (
	"log"

	"github.com/hidenari-yuda/go-grpc-clean/pb"
	"github.com/hidenari-yuda/go-grpc-clean/usecase"
)

type MediaInteractor interface {
	// Gest API
	// Create
	Create(user *pb.Media) (*pb.Media, error)

	// Update
	Update(user *pb.Media) (bool, error)

	// Delete
	Delete(id int64) (bool, error)

	// Get
	GetById(id int64) (*pb.Media, error)

	// get list by user id
	GetListByUserId(userId int64) ([]*pb.Media, error)

	// get list by type
	GetListByType(mediaType int64) ([]*pb.Media, error)

	// admin API
	GetAll() ([]*pb.Media, error)
}

type MediaInteractorImpl struct {
	firebase       usecase.Firebase
	mediaRepository usecase.MediaRepository
}

func NewMediaInteractorImpl(
	fb usecase.Firebase,
	uR usecase.MediaRepository,
) MediaInteractor {
	return &MediaInteractorImpl{
		firebase:       fb,
		mediaRepository: uR,
	}
}

// Create
func (i *MediaInteractorImpl) Create(user *pb.Media) (*pb.Media, error) {

	// ユーザー登録
	err := i.mediaRepository.Create(user)
	if err != nil {
		return user, err
	}

	return user, nil
}

// Update
func (i *MediaInteractorImpl) Update(user *pb.Media) (bool, error) {

	// ユーザー登録
	err := i.mediaRepository.Update(user)
	if err != nil {
		return false, err
	}

	return true, nil
}

// delete
func (i *MediaInteractorImpl) Delete(id int64) (bool, error) {

	err := i.mediaRepository.Delete(id)
	if err != nil {
		return false, err
	}

	return true, nil
}

// GetById
func (i *MediaInteractorImpl) GetById(id int64) (*pb.Media, error) {
	var (
		user *pb.Media
		err  error
	)

	// ユーザー登録
	user, err = i.mediaRepository.GetById(id)
	if err != nil {
		log.Println("error is:", err)
		return user, err
	}

	return user, nil
}

// GetListByUserId
func (i *MediaInteractorImpl) GetListByUserId(userId int64) ([]*pb.Media, error) {
	var (
		users []*pb.Media
		err   error
	)

	users, err = i.mediaRepository.GetListByUserId(userId)
	if err != nil {
		return users, err
	}

	return users, nil
}

// GetListByType
func (i *MediaInteractorImpl) GetListByType(mediaType int64) ([]*pb.Media, error) {
	var (
		users []*pb.Media
		err   error
	)

	users, err = i.mediaRepository.GetListByType(mediaType)
	if err != nil {
		return users, err
	}

	return users, nil
}


// GetAll
func (i *MediaInteractorImpl) GetAll() ([]*pb.Media, error) {
	var (
		users []*pb.Media
		err   error
	)

	users, err = i.mediaRepository.GetAll()
	if err != nil {
		return users, err
	}

	// i.userClient.DetectTextFromImage()

	return users, nil
}
