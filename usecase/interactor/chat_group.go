package interactor

import (
	"fmt"

	"github.com/hidenari-yuda/go-grpc-clean/pb"
	"github.com/hidenari-yuda/go-grpc-clean/usecase"
)

type ChatGroupInteractor interface {
	// Gest API
	// Create
	Create(ChatGroup *pb.ChatGroup) (*pb.ChatGroup, error)

	// Update
	Update(ChatGroup *pb.ChatGroup) (bool, error)

	// Delete
	Delete(id uint) (bool, error)

	// Get
	GetById(id uint) (*pb.ChatGroup, error)

	// get list by user id
	GetListByUserId(userId uint) ([]*pb.ChatGroup, error)
}

type ChatGroupInteractorImpl struct {
	firebase usecase.Firebase
	chatGroupRepository usecase.ChatGroupRepository
	chatUserRepository	usecase.ChatUserRepository
}

func NewChatGroupInteractorImpl(
	fb usecase.Firebase,
	cgR usecase.ChatGroupRepository,
	cuR usecase.ChatUserRepository,
) ChatGroupInteractor {
	return &ChatGroupInteractorImpl{
		firebase: fb,
		chatGroupRepository: cgR,
		chatUserRepository:  cuR,
	}
}

func (i *ChatGroupInteractorImpl) Create(chatGroup *pb.ChatGroup) (*pb.ChatGroup, error) {

	// ユーザー登録
	err := i.chatGroupRepository.Create(chatGroup)
	if err != nil {
		return chatGroup, err
	}

	return chatGroup, nil
}

func (i *ChatGroupInteractorImpl) Update(chatGroup *pb.ChatGroup) (bool, error) {

	// ユーザー登録
	err := i.chatGroupRepository.Update(chatGroup)
	if err != nil {
		return false, err
	}

	return true, nil
}

func (i *ChatGroupInteractorImpl) Delete(id uint) (bool, error) {

	// ユーザー登録
	err := i.chatGroupRepository.Delete(id)
	if err != nil {
		return false, err
	}

	return true, nil
}

func (i *ChatGroupInteractorImpl) GetById(id uint) (*pb.ChatGroup, error) {
	var (
		chatGroup *pb.ChatGroup
		err       error
	)

	// ユーザー登録
	chatGroup, err = i.chatGroupRepository.GetById(id)
	if err != nil {
		fmt.Println("error is:", err)
		return chatGroup, err
	}

	return chatGroup, nil
}

func (i *ChatGroupInteractorImpl) GetListByUserId(userId uint) ([]*pb.ChatGroup, error) {
	var (
		chatGroups []*pb.ChatGroup
		err        error
	)

	// ユーザー登録
	chatGroups, err = i.chatGroupRepository.GetListByUserId(userId)
	if err != nil {
		fmt.Println("error is:", err)
		return chatGroups, err
	}

	return chatGroups, nil
}