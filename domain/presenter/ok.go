package presenter

import "github.com/hidenari-yuda/go-grpc-clean/domain/responses"

func NewOkJSONPresenter(resp responses.OK) Presenter {
	return NewJSONPresenter(200, resp)
}
