package presenter

import (
	"demo/go-clean-demo/presenter/viewinterface"
	"demo/go-clean-demo/presenter/viewmodel"
	"demo/go-clean-demo/usecase/ucoutput"
)

type ErrorPresenter struct {
	v viewinterface.ResponseSender
}

func InitErrorPresenter(v viewinterface.ResponseSender) *ErrorPresenter {
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
