package presenter

import "github.com/hidenari-yuda/go-docker-template/domain/entity/responses"

func NewOkJSONPresenter(resp responses.OK) Presenter {
	return NewJSONPresenter(200, resp)
}
