package presenter

import (
	"demo/go-clean-demo/presenter/viewmodel"
	"demo/go-clean-demo/usecase/ucoutput"
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

func (p *ErrorPresenter) PresentErrorResponse(err ucoutput.Error) {
	p.v.SendResponse(err.ErrorStatus, viewmodel.ErrorMessage{
		ErrorCode: err.ErrorCode,
		ErrorMsg:  err.ErrorMessage,
	})
}
