package interactor

import (
	"log"

	"github.com/hidenari-yuda/go-grpc-clean/pb"
	"github.com/hidenari-yuda/go-grpc-clean/usecase"
)

type KeywordInteractor interface {
	// Gest API
	// Create
	Create(Keyword *pb.Keyword) (*pb.Keyword, error)

	// Update
	Update(Keyword *pb.Keyword) (bool, error)

	// Delete
	Delete(id int64) (bool, error)

	// Get
	GetById(id int64) (*pb.Keyword, error)

	// get list by user id
	GetListByUserId(userId int64) ([]*pb.Keyword, error)
}

type KeywordInteractorImpl struct {
	firebase            usecase.Firebase
	keywordRepository usecase.KeywordRepository
	chatUserRepository  usecase.ChatUserRepository
}

func NewKeywordInteractorImpl(
	fb usecase.Firebase,
	cgR usecase.KeywordRepository,
	cuR usecase.ChatUserRepository,
) KeywordInteractor {
	return &KeywordInteractorImpl{
		firebase:            fb,
		keywordRepository: cgR,
		chatUserRepository:  cuR,
	}
}

func (i *KeywordInteractorImpl) Create(keyword *pb.Keyword) (*pb.Keyword, error) {

	// ユーザー登録
	err := i.keywordRepository.Create(keyword)
	if err != nil {
		return keyword, err
	}

	return keyword, nil
}

func (i *KeywordInteractorImpl) Update(keyword *pb.Keyword) (bool, error) {

	// ユーザー登録
	err := i.keywordRepository.Update(keyword)
	if err != nil {
		return false, err
	}

	return true, nil
}

func (i *KeywordInteractorImpl) Delete(id int64) (bool, error) {

	// ユーザー登録
	err := i.keywordRepository.Delete(id)
	if err != nil {
		return false, err
	}

	return true, nil
}

func (i *KeywordInteractorImpl) GetById(id int64) (*pb.Keyword, error) {
	var (
		keyword *pb.Keyword
		err       error
	)

	// ユーザー登録
	keyword, err = i.keywordRepository.GetById(id)
	if err != nil {
		log.Println("error is:", err)
		return keyword, err
	}

	return keyword, nil
}

func (i *KeywordInteractorImpl) GetListByUserId(userId int64) ([]*pb.Keyword, error) {
	var (
		keywords []*pb.Keyword
		err        error
	)

	// ユーザー登録
	// keywords, err = i.keywordRepository.GetListByUserId(userId)
	if err != nil {
		log.Println("error is:", err)
		return keywords, err
	}

	return keywords, nil
}
