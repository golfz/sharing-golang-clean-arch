package presenter

import (
	"demo/go-clean-demo/presenter/viewmodel"
	"demo/go-clean-demo/view"
)

type ErrorPresenter struct {
	v *view.JsonResponseView
}

func InitErrorPresenter(v *view.JsonResponseView) *ErrorPresenter {
	return &ErrorPresenter{
		v: v,
	}
}

func (p *ErrorPresenter) PresentErrorResponse(status int, errBody viewmodel.ErrorMessage) {
	p.v.SendResponse(status, errBody)
}
