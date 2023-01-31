package interactor

import (
	"fmt"

	"github.com/hidenari-yuda/go-grpc-clean/domain/entity"
	"github.com/hidenari-yuda/go-grpc-clean/usecase"
)

type ChatInteractor interface {
	// Gest API
	// Create
	Create(Chat *entity.Chat) (*entity.Chat, error)

	// Update
	// Update(Chat *entity.Chat) (*entity.Chat, error)

	// Get
	GetById(id uint) (*entity.Chat, error)
	GetStream(groupId uint) ([]*entity.Chat, error)
}

type ChatInteractorImpl struct {
	firebase            usecase.Firebase
	chatRepository      usecase.ChatRepository
	chatGroupRepository usecase.ChatGroupRepository
	chatUserRepository  usecase.ChatUserRepository
}

func NewChatInteractorImpl(
	fb usecase.Firebase,
	cR usecase.ChatRepository,
	cgR usecase.ChatGroupRepository,
	cuR usecase.ChatUserRepository,
) ChatInteractor {
	return &ChatInteractorImpl{
		firebase:            fb,
		chatRepository:      cR,
		chatGroupRepository: cgR,
		chatUserRepository:  cuR,
	}
}

func (i *ChatInteractorImpl) Create(chat *entity.Chat) (*entity.Chat, error) {
	var (
		err error
	)

	if chat.GroupId == 0 {
	}

	// ユーザー登録
	err = i.chatRepository.Create(chat)
	if err != nil {
		return chat, err
	}

	return chat, nil
}

// func (i *ChatInteractorImpl) Update(chat *entity.Chat) (*entity.Chat, error) {
// 	var (
// 		err error
// 	)

// 	// ユーザー登録
// 	err = i.ChatRepository.Create(chat)
// 	if err != nil {
// 		return chat, err
// 	}

// 	return chat, nil
// }

func (i *ChatInteractorImpl) GetById(id uint) (*entity.Chat, error) {
	var (
		chat *entity.Chat
		err  error
	)

	// ユーザー登録
	chat, err = i.chatRepository.GetById(id)
	if err != nil {
		fmt.Println("error is:", err)
		return chat, err
	}

	return chat, nil
}

func (i *ChatInteractorImpl) GetStream(groupId uint) ([]*entity.Chat, error) {
	var (
		chats []*entity.Chat
		err   error
	)

	// ユーザー登録
	chats, err = i.chatRepository.GetListByGroupId(groupId)
	if err != nil {
		fmt.Println("error is:", err)
		return chats, err
	}

	return chats, nil
}
