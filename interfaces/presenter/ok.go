package presenter

import "github.com/hidenari-yuda/umerun-resume/domain/entity/responses"

func NewOkJSONPresenter(resp responses.OK) Presenter {
	return NewJSONPresenter(200, resp)
}
