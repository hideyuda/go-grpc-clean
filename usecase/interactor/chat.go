package interactor

import (
	"context"
	"fmt"
	"log"

	"github.com/hidenari-yuda/go-grpc-clean/pb"
	"github.com/hidenari-yuda/go-grpc-clean/usecase"
	"golang.org/x/sync/errgroup"
)

type ChatInteractor interface {
	// Gest API
	// Create
	Create(Chat *pb.Chat) (*pb.Chat, error)

	// Update
	Update(Chat *pb.Chat) (bool, error)

	// Delete
	Delete(id uint) (bool, error)

	// Get
	GetById(id uint) (*pb.Chat, error)

	// get list by user id
	GetListByGroupId(groupId uint) ([]*pb.Chat, error)

	// stream
	GetStream(ctx context.Context, stream chan<- pb.Chat) error
}

type ChatInteractorImpl struct {
	firebase usecase.Firebase
	// chatRepository      usecase.ChatRepository
	chatRepository usecase.ChatRepository
	// chatUserRepository  usecase.ChatRepository
}

func NewChatInteractorImpl(
	fb usecase.Firebase,
	// cR usecase.ChatRepository,
	cgR usecase.ChatRepository,
	// cuR usecase.ChatUserRepository,
) ChatInteractor {
	return &ChatInteractorImpl{
		firebase: fb,
		// chatRepository:      cR,
		chatRepository: cgR,
		// chatUserRepository:  cuR,
	}
}

func (i *ChatInteractorImpl) Create(chat *pb.Chat) (*pb.Chat, error) {

	// ユーザー登録
	err := i.chatRepository.Create(chat)
	if err != nil {
		return chat, err
	}

	return chat, nil
}

func (i *ChatInteractorImpl) Update(chat *pb.Chat) (bool, error) {

	// ユーザー登録
	err := i.chatRepository.Update(chat)
	if err != nil {
		return false, err
	}

	return true, nil
}

func (i *ChatInteractorImpl) Delete(id uint) (bool, error) {

	// ユーザー登録
	err := i.chatRepository.Delete(id)
	if err != nil {
		return false, err
	}

	return true, nil
}

func (i *ChatInteractorImpl) GetById(id uint) (*pb.Chat, error) {
	var (
		chat *pb.Chat
		err  error
	)

	// ユーザー登録
	chat, err = i.chatRepository.GetById(id)
	if err != nil {
		log.Println("error is:", err)
		return chat, err
	}

	return chat, nil
}

func (i *ChatInteractorImpl) GetListByGroupId(groupId uint) ([]*pb.Chat, error) {
	var (
		chats []*pb.Chat
		err   error
	)

	// ユーザー登録
	chats, err = i.chatRepository.GetListByGroupId(groupId)
	if err != nil {
		log.Println("error is:", err)
		return chats, err
	}

	return chats, nil
}

// get stream
func (i *ChatInteractorImpl) GetStream(ctx context.Context, stream chan<- pb.Chat) error {
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
