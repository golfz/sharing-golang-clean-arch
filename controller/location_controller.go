package controller

import (
	"demo/go-clean-demo/controller/request"
	"demo/go-clean-demo/usecase/interface/pinterface"
	"demo/go-clean-demo/usecase/interface/ucinterface"
	"demo/go-clean-demo/usecase/ucinput"
	"demo/go-clean-demo/usecase/ucoutput"
	"encoding/json"
	"net/http"
	"time"
)

type LocationController struct {
	request  *http.Request
	pSuccess pinterface.AddLocationResponsePresenter
	pError   pinterface.ErrorResponsePresenter
}

func InitLocationController(r *http.Request, pSuccess pinterface.AddLocationResponsePresenter, pError pinterface.ErrorResponsePresenter) *LocationController {
	return &LocationController{
		request:  r,
		pSuccess: pSuccess,
		pError:   pError,
	}
}

func (ctrl *LocationController) AddLocation(uc ucinterface.LocationAdder) {
	reqData := request.AddLocation{}

	errReqData := json.NewDecoder(ctrl.request.Body).Decode(&reqData)
	if errReqData != nil {

		ctrl.pError.PresentErrorResponse(ucoutput.Error{
			ErrorStatus:  http.StatusBadRequest,
			ErrorCode:    http.StatusBadRequest,
			ErrorMessage: "request body mismatched",
		})

		return
	}

	t, errTime := time.Parse("2006-01-02 15:04:05Z07:00", reqData.Datetime)
	if errTime != nil {
		ctrl.pError.PresentErrorResponse(ucoutput.Error{
			ErrorStatus:  http.StatusBadRequest,
			ErrorCode:    http.StatusBadRequest,
			ErrorMessage: "time-format mismatched",
		})
		return
	}

	if (reqData.Lat < -90 || 90 < reqData.Lat) || (reqData.Long < -180 || 180 < reqData.Long) {
		ctrl.pError.PresentErrorResponse(ucoutput.Error{
			ErrorStatus:  http.StatusBadRequest,
			ErrorCode:    http.StatusBadRequest,
			ErrorMessage: "Lat or Long is not corrected",
		})
		return
	}

	useCaseInput := ucinput.NewLocation{
		Time: t,
		Lat:  reqData.Lat,
		Long: reqData.Long,
	}

	uc.AddLocation(useCaseInput, ctrl.pSuccess)
}
