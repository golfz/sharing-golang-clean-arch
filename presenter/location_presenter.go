package presenter

import (
	"demo/go-clean-demo/entity"
	"demo/go-clean-demo/presenter/viewmodel"
	"demo/go-clean-demo/view"
	"fmt"
	"net/http"
	"time"
)

type LocationPresenter struct {
	v *view.JsonResponseView
}

func InitLocationPresenter(v *view.JsonResponseView) *LocationPresenter {
	return &LocationPresenter{
		v: v,
	}
}

func (p *LocationPresenter) PresentAddLocationResponse(responseBody []entity.Location) {

	resp := []viewmodel.LocationData{}

	for _, v := range responseBody {
		kmh := float64(v.GetSpeedMPH()) * 1.60934
		speed := fmt.Sprintf("%d km/h", int64(kmh))

		resp = append(resp, viewmodel.LocationData{
			Id:       *v.Id,
			Datetime: v.Time.Format(time.RFC1123),
			Lat:      v.Lat,
			Long:     v.Long,
			Speed:    speed,
		})
	}

	p.v.SendResponse(http.StatusCreated, resp)
}
