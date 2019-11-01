package presenter

import (
	"demo/go-clean-demo/entity"
	"demo/go-clean-demo/view"
	"fmt"
	"net/http"
	"time"
)

type responseData struct {
	Id       int64   `json:"id"`
	Datetime string  `json:"datetime"`
	Lat      float64 `json:"lat"`
	Long     float64 `json:"long"`
	Speed    string  `json:"speed"`
}

type ErrorMessage struct {
	ErrorCode int    `json:"error_code"`
	ErrorMsg  string `json:"error_msg"`
}

type LocationPresenter struct {
	v *view.JsonResponseView
	w http.ResponseWriter
}

func InitLocationPresenter(w http.ResponseWriter, v *view.JsonResponseView) *LocationPresenter {
	return &LocationPresenter{
		w: w,
		v: v,
	}
}

func (p *LocationPresenter) PresentAddLocationResponse(responseBody []entity.Location) {

	resp := []responseData{}

	for _, v := range responseBody {
		kmh := float64(v.GetSpeedMPH()) * 1.60934
		speed := fmt.Sprintf("%d km/h", int64(kmh))

		resp = append(resp, responseData{
			Id:       *v.Id,
			Datetime: v.Time.Format(time.RFC1123),
			Lat:      v.Lat,
			Long:     v.Long,
			Speed:    speed,
		})
	}

	p.v.SendResponse(http.StatusInternalServerError, resp)
}
