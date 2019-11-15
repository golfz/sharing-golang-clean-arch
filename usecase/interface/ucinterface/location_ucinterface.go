package ucinterface

import (
	"demo/go-clean-demo/usecase/interface/pinterface"
	"demo/go-clean-demo/usecase/ucinput"
)

type LocationAdder interface {
	AddLocation(inputData ucinput.NewLocation, pSuccess pinterface.AddLocationResponsePresenter)
}
