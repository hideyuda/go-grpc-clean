package presenter

import "github.com/hidenari-yuda/go-grpc-clean/domain/responses"

func NewChatJSONPresenter(resp responses.Chat) Presenter {
	return NewJSONPresenter(200, resp)
}

func NewChatListJSONPresenter(resp responses.ChatList) Presenter {
	return NewJSONPresenter(200, resp)
}
