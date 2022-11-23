package presenter

import "github.com/hidenari-yuda/umerun-resume/domain/entity/responses"

func NewUserJSONPresenter(resp responses.User) Presenter {
	return NewJSONPresenter(200, resp)
}

func NewUserListJSONPresenter(resp responses.UserList) Presenter {
	return NewJSONPresenter(200, resp)
}
