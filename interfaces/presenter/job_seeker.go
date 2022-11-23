package presenter

import (
	"github.com/hidenari-yuda/umerun-resume/domain/entity/responses"
)

func NewJobSeekerJSONPresenter(resp responses.JobSeeker) Presenter {
	return NewJSONPresenter(200, resp)
}

func NewJobSeekerListJSONPresenter(resp responses.JobSeekerList) Presenter {
	return NewJSONPresenter(200, resp)
}
