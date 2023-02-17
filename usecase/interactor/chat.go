package interactor

import (
	"context"
	"fmt"

	"github.com/hidenari-yuda/go-grpc-clean/domain/entity"
	"github.com/hidenari-yuda/go-grpc-clean/usecase"
	"golang.org/x/sync/errgroup"
)

type ChatInteractor interface {
	// Gest API
	// Create
	Create(Chat *entity.Chat) (*entity.Chat, error)

	// Update
	// Update(Chat *entity.Chat) (*entity.Chat, error)

	// Get
	GetById(id uint) (*entity.Chat, error)
	GetStream(ctx context.Context, stream chan<- entity.Chat) error
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

	// ユーザー登録
	err := i.chatRepository.Create(chat)
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

func (i *ChatInteractorImpl) GetStream(ctx context.Context, stream chan<- entity.Chat) error {
	defer close(stream)
	eg, _ := errgroup.WithContext(ctx)
	eg.Go(func() error {
		err := i.firebase.GetChatStream(ctx, stream)
		if err != nil {
			return err
		}
		return nil
	})
	if err := eg.Wait(); err != nil {
		return fmt.Errorf("failed to GetMessageStreamService.Handle: %s", err)
	}
	return nil
}
