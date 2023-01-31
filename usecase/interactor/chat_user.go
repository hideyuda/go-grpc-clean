package interactor

import (
	"fmt"

	"github.com/hidenari-yuda/go-grpc-clean/domain/entity"
	"github.com/hidenari-yuda/go-grpc-clean/usecase"
)

type ChatUserInteractor interface {
	// Gest API
	// Create
	Create(ChatUser *entity.ChatUser) (*entity.ChatUser, error)

	// Update
	// Update(ChatUser *entity.ChatUser) (*entity.ChatUser, error)

	// Get
	GetById(id uint) (*entity.ChatUser, error)
	GetListByGroupId(groupId uint) ([]*entity.ChatUser, error)
}

type ChatUserInteractorImpl struct {
	firebase usecase.Firebase
	// chatUserRepository  usecase.ChatUserRepository
	chatGroupRepository usecase.ChatGroupRepository
	chatUserRepository  usecase.ChatUserRepository
}

func NewChatUserInteractorImpl(
	fb usecase.Firebase,
	// cR usecase.ChatUserRepository,
	cgR usecase.ChatGroupRepository,
	cuR usecase.ChatUserRepository,
) ChatUserInteractor {
	return &ChatUserInteractorImpl{
		firebase: fb,
		// chatUserRepository:      cR,
		chatGroupRepository: cgR,
		chatUserRepository:  cuR,
	}
}

func (i *ChatUserInteractorImpl) Create(chatUser *entity.ChatUser) (*entity.ChatUser, error) {
	var (
		err error
	)

	if chatUser.GroupId == 0 {
	}

	// ユーザー登録
	err = i.chatUserRepository.Create(chatUser)
	if err != nil {
		return chatUser, err
	}

	return chatUser, nil
}

// func (i *ChatUserInteractorImpl) Update(chatUser *entity.ChatUser) (*entity.ChatUser, error) {
// 	var (
// 		err error
// 	)

// 	// ユーザー登録
// 	err = i.ChatUserRepository.Create(chatUser)
// 	if err != nil {
// 		return chatUser, err
// 	}

// 	return chatUser, nil
// }

func (i *ChatUserInteractorImpl) GetById(id uint) (*entity.ChatUser, error) {
	var (
		chatUser *entity.ChatUser
		err      error
	)

	// ユーザー登録
	chatUser, err = i.chatUserRepository.GetById(id)
	if err != nil {
		fmt.Println("error is:", err)
		return chatUser, err
	}

	return chatUser, nil
}

func (i *ChatUserInteractorImpl) GetListByGroupId(groupId uint) ([]*entity.ChatUser, error) {
	var (
		chatUsers []*entity.ChatUser
		err       error
	)

	// ユーザー登録
	chatUsers, err = i.chatUserRepository.GetListByGroupId(groupId)
	if err != nil {
		fmt.Println("error is:", err)
		return chatUsers, err
	}

	return chatUsers, nil
}
