package handler

import "github.com/hidenari-yuda/go-grpc-clean/usecase/interactor"

type UserHandler interface {
	// Gest API
	// Create(param *entity.User) (presenter.Presenter, error)
	// Update(param *entity.User) (presenter.Presenter, error)

	// // Get
	// GetById(id uint) (presenter.Presenter, error)
	// SignIn(param *entity.SignInParam) (presenter.Presenter, error)
}

type UserHandlerImpl struct {
	UserInteractor interactor.UserInteractor
}

func NewUserHandlerImpl(ui interactor.UserInteractor) UserHandler {
	return &UserHandlerImpl{
		UserInteractor: ui,
	}
}

// func (h *UserHandlerImpl) Create(param *entity.User) (presenter.Presenter, error) {
// 	output, err := h.UserInteractor.Create(param)
// 	if err != nil {
// 		// c.JSON(c, presenter.NewErrorJsonPresenter(err))
// 		return nil, err
// 	}

// 	return presenter.NewUserJSONPresenter(responses.NewUser(output)), nil
// }

// func (h *UserHandlerImpl) Update(param *entity.User) (presenter.Presenter, error) {
// 	output, err := h.UserInteractor.Update(param)
// 	if err != nil {
// 		// c.JSON(c, presenter.NewErrorJsonPresenter(err))
// 		return nil, err
// 	}

// 	return presenter.NewUserJSONPresenter(responses.NewUser(output)), nil
// }

// // Get
// func (h *UserHandlerImpl) GetById(id uint) (presenter.Presenter, error) {
// 	output, err := h.UserInteractor.GetById(id)
// 	if err != nil {
// 		// c.JSON(c, presenter.NewErrorJsonPresenter(err))
// 		return nil, err
// 	}

// 	return presenter.NewUserJSONPresenter(responses.NewUser(output)), nil
// }

// func (h *UserHandlerImpl) SignIn(param *entity.SignInParam) (presenter.Presenter, error) {
// 	output, err := h.UserInteractor.SignIn(param)
// 	if err != nil {
// 		// c.JSON(c, presenter.NewErrorJsonPresenter(err))
// 		return nil, err
// 	}

// 	return presenter.NewUserJSONPresenter(responses.NewUser(output)), nil

// }
