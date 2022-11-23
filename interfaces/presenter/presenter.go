package presenter

type Presenter interface {
	StatusCode() int
	Data() interface{}
}

type PresenterImpl struct {
	statusCode int
	data       interface{}
}

func (p *PresenterImpl) StatusCode() int {
	return p.statusCode
}

func (p *PresenterImpl) Data() interface{} {
	return p.data
}

type JSONPresenter struct {
	PresenterImpl
}

func NewJSONPresenter(statusCode int, data interface{}) *JSONPresenter {
	return &JSONPresenter{PresenterImpl{statusCode, data}}
}
