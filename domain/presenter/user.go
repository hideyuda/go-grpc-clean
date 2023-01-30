package presenter

import "github.com/hidenari-yuda/go-grpc-clean/domain/responses"

func NewUserJSONPresenter(resp responses.User) Presenter {
	return NewJSONPresenter(200, resp)
}

func NewUserListJSONPresenter(resp responses.UserList) Presenter {
	return NewJSONPresenter(200, resp)
}
