package controller

import (
	"demo/go-clean-demo/presenter"
	"demo/go-clean-demo/usecase"
	"demo/go-clean-demo/usecase/ucinput"
	"demo/go-clean-demo/usecase/ucoutput"
	"encoding/json"
	"net/http"
	"time"
)

type requestData struct {
	Datetime string  `json:"datetime"`
	Lat      float64 `json:"lat"`
	Long     float64 `json:"long"`
}

type LocationCtrl struct {
	request  *http.Request
	pSuccess *presenter.LocationPresenter
	pError   *presenter.ErrorPresenter
}

func InitLocationController(r *http.Request, pSuccess *presenter.LocationPresenter, pError *presenter.ErrorPresenter) *LocationCtrl {
	return &LocationCtrl{
		request:  r,
		pSuccess: pSuccess,
		pError:   pError,
	}
}

func (ctrl *LocationCtrl) AddLocation(uc *usecase.LocationUseCase) {
	reqData := requestData{}

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
