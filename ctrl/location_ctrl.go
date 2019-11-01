package ctrl

import (
	"demo/go-clean-demo/entity"
	"demo/go-clean-demo/fakedb"
	"demo/go-clean-demo/model"
	"demo/go-clean-demo/presenter"
	"demo/go-clean-demo/view"
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
	writer   http.ResponseWriter
	request  *http.Request
	v        *view.JsonResponseView
	pSuccess *presenter.LocationPresenter
}

func InitLocationCtrl(w http.ResponseWriter, r *http.Request, v *view.JsonResponseView, pSuccess *presenter.LocationPresenter) *LocationCtrl {
	return &LocationCtrl{
		writer:   w,
		request:  r,
		v:        v,
		pSuccess: pSuccess,
	}
}

func (ctrl *LocationCtrl) AddLocationCtrl() {
	reqData := requestData{}

	errReqData := json.NewDecoder(ctrl.request.Body).Decode(&reqData)
	if errReqData != nil {

		ctrl.v.SendResponse(http.StatusInternalServerError, presenter.ErrorMessage{
			ErrorCode: http.StatusBadRequest,
			ErrorMsg:  "request body mismatched",
		})
		return
	}

	t, errTime := time.Parse("2006-01-02 15:04:05Z07:00", reqData.Datetime)
	if errTime != nil {
		ctrl.v.SendResponse(http.StatusInternalServerError, presenter.ErrorMessage{
			ErrorCode: http.StatusBadRequest,
			ErrorMsg:  "time-format mismatched",
		})
		return
	}

	if (reqData.Lat < -90 || 90 < reqData.Lat) || (reqData.Long < -180 || 180 < reqData.Long) {
		ctrl.v.SendResponse(http.StatusInternalServerError, presenter.ErrorMessage{
			ErrorCode: http.StatusBadRequest,
			ErrorMsg:  "Lat or Long is not corrected",
		})
		return
	}

	db := fakedb.InitDBConnection()

	locationModel := model.InitLocationModel(db);
	locationModel.AddNewLocation(entity.Location{
		Time: t,
		Lat:  reqData.Lat,
		Long: reqData.Long,
	})

	locationList := locationModel.GetAll()

	ctrl.pSuccess.PresentAddLocationResponse(locationList)

	return
}
