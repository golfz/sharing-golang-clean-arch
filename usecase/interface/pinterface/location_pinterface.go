package pinterface

import "demo/go-clean-demo/usecase/ucoutput"

type AddLocationResponsePresenter interface {
	PresentAddLocationResponse(responseBody []ucoutput.Location)
}
