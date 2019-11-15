package pinterface

import "demo/go-clean-demo/usecase/ucoutput"

type ErrorResponsePresenter interface {
	PresentErrorResponse(err ucoutput.Error)
}
