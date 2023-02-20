package interactor

// import (
// 	"fmt"

// 	"github.com/hidenari-yuda/go-grpc-clean/domain/entity"
// 	"github.com/hidenari-yuda/go-grpc-clean/usecase"
// )

// type ChatGroupInteractor interface {
// 	// Gest API
// 	// Create
// 	Create(ChatGroup *entity.ChatGroup) (*entity.ChatGroup, error)

// 	// Update
// 	Update(ChatGroup *entity.ChatGroup) (*entity.ChatGroup, error)

// 	// Get
// 	GetById(id uint) (*entity.ChatGroup, error)
// }

// type ChatGroupInteractorImpl struct {
// 	firebase usecase.Firebase
// 	// chatGroupRepository      usecase.ChatGroupRepository
// 	chatGroupRepository usecase.ChatGroupRepository
// 	// chatGroupUserRepository  usecase.ChatGroupRepository
// }

// func NewChatGroupInteractorImpl(
// 	fb usecase.Firebase,
// 	// cR usecase.ChatGroupRepository,
// 	cgR usecase.ChatGroupRepository,
// 	// cuR usecase.ChatUserRepository,
// ) ChatGroupInteractor {
// 	return &ChatGroupInteractorImpl{
// 		firebase: fb,
// 		// chatRepository:      cR,
// 		chatGroupRepository: cgR,
// 		// chatUserRepository:  cuR,
// 	}
// }

// func (i *ChatGroupInteractorImpl) Create(chatGroup *entity.ChatGroup) (*entity.ChatGroup, error) {

// 	// ユーザー登録
// 	err := i.chatGroupRepository.Create(chatGroup)
// 	if err != nil {
// 		return chatGroup, err
// 	}

// 	return chatGroup, nil
// }

// func (i *ChatGroupInteractorImpl) Update(chatGroup *entity.ChatGroup) (*entity.ChatGroup, error) {

// 	// ユーザー登録
// 	err := i.chatGroupRepository.Update(chatGroup)
// 	if err != nil {
// 		return chatGroup, err
// 	}

// 	return chatGroup, nil
// }

// func (i *ChatGroupInteractorImpl) GetById(id uint) (*entity.ChatGroup, error) {
// 	var (
// 		chatGroup *entity.ChatGroup
// 		err       error
// 	)

// 	// ユーザー登録
// 	chatGroup, err = i.chatGroupRepository.GetById(id)
// 	if err != nil {
// 		fmt.Println("error is:", err)
// 		return chatGroup, err
// 	}

// 	return chatGroup, nil
// }
