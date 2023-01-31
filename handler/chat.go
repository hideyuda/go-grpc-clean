package handler

import (
	"github.com/hidenari-yuda/go-grpc-clean/domain/entity"
	"github.com/hidenari-yuda/go-grpc-clean/domain/presenter"
	"github.com/hidenari-yuda/go-grpc-clean/domain/responses"

	"github.com/hidenari-yuda/go-grpc-clean/usecase/interactor"
)

type ChatHandler interface {
	// Gest API
	Create(param *entity.Chat) (presenter.Presenter, error)
	// Update(param *entity.Chat) (presenter.Presenter, error)

	// Get
	GetById(id uint) (presenter.Presenter, error)
	GetStream(groupId uint) (presenter.Presenter, error)
}

type ChatHandlerImpl struct {
	ChatInteractor interactor.ChatInteractor
}

func NewChatHandlerImpl(ui interactor.ChatInteractor) ChatHandler {
	return &ChatHandlerImpl{
		ChatInteractor: ui,
	}
}

func (h *ChatHandlerImpl) Create(param *entity.Chat) (presenter.Presenter, error) {
	output, err := h.ChatInteractor.Create(param)
	if err != nil {
		// c.JSON(c, presenter.NewErrorJsonPresenter(err))
		return nil, err
	}

	return presenter.NewChatJSONPresenter(responses.NewChat(output)), nil
}

// func (h *ChatHandlerImpl) Update(param *entity.Chat) (presenter.Presenter, error) {
// 	output, err := h.ChatInteractor.Update(param)
// 	if err != nil {
// 		// c.JSON(c, presenter.NewErrorJsonPresenter(err))
// 		return nil, err
// 	}

// 	return presenter.NewChatJSONPresenter(responses.NewChat(output)), nil
// }

// Get
func (h *ChatHandlerImpl) GetById(id uint) (presenter.Presenter, error) {
	output, err := h.ChatInteractor.GetById(id)
	if err != nil {
		// c.JSON(c, presenter.NewErrorJsonPresenter(err))
		return nil, err
	}

	return presenter.NewChatJSONPresenter(responses.NewChat(output)), nil
}

func (h *ChatHandlerImpl) GetStream(groupId uint) (presenter.Presenter, error) {
	output, err := h.ChatInteractor.GetStream(groupId)
	if err != nil {
		// c.JSON(c, presenter.NewErrorJsonPresenter(err))
		return nil, err
	}

	return presenter.NewChatListJSONPresenter(responses.NewChatList(output)), nil
}
