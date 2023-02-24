package interactor

import (
	"log"

	"github.com/hidenari-yuda/go-grpc-clean/pb"
	"github.com/hidenari-yuda/go-grpc-clean/usecase"
)

type ScriptInteractor interface {
	// Gest API
	// Create
	Create(user *pb.Script) (*pb.Script, error)

	// Update
	Update(user *pb.Script) (bool, error)

	// Delete
	Delete(id int64) (bool, error)

	// Get
	GetById(id int64) (*pb.Script, error)

	// get list by media id
	GetListByMediaId(mediaId int64) ([]*pb.Script, error)

	// admin API
	GetAll() ([]*pb.Script, error)
}

type ScriptInteractorImpl struct {
	firebase       usecase.Firebase
	scriptRepository usecase.ScriptRepository
}

func NewScriptInteractorImpl(
	fb usecase.Firebase,
	uR usecase.ScriptRepository,
) ScriptInteractor {
	return &ScriptInteractorImpl{
		firebase:       fb,
		scriptRepository: uR,
	}
}

// Create
func (i *ScriptInteractorImpl) Create(user *pb.Script) (*pb.Script, error) {

	// ユーザー登録
	err := i.scriptRepository.Create(user)
	if err != nil {
		return user, err
	}

	return user, nil
}

// Update
func (i *ScriptInteractorImpl) Update(user *pb.Script) (bool, error) {

	// ユーザー登録
	err := i.scriptRepository.Update(user)
	if err != nil {
		return false, err
	}

	return true, nil
}

// delete
func (i *ScriptInteractorImpl) Delete(id int64) (bool, error) {

	err := i.scriptRepository.Delete(id)
	if err != nil {
		return false, err
	}

	return true, nil
}

// GetById
func (i *ScriptInteractorImpl) GetById(id int64) (*pb.Script, error) {
	var (
		user *pb.Script
		err  error
	)

	// ユーザー登録
	user, err = i.scriptRepository.GetById(id)
	if err != nil {
		log.Println("error is:", err)
		return user, err
	}

	return user, nil
}

// GetListByMediaId
func (i *ScriptInteractorImpl) GetListByMediaId(userId int64) ([]*pb.Script, error) {
	var (
		users []*pb.Script
		err   error
	)

	users, err = i.scriptRepository.GetListByMediaId(userId)
	if err != nil {
		return users, err
	}

	return users, nil
}


// GetAll
func (i *ScriptInteractorImpl) GetAll() ([]*pb.Script, error) {
	var (
		users []*pb.Script
		err   error
	)

	users, err = i.scriptRepository.GetAll()
	if err != nil {
		return users, err
	}

	// i.userClient.DetectTextFromImage()

	return users, nil
}
