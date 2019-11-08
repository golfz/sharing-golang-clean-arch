package usecase

import (
	"demo/go-clean-demo/entity"
	"demo/go-clean-demo/fakedb"
	"demo/go-clean-demo/model"
	"demo/go-clean-demo/presenter"
	"demo/go-clean-demo/usecase/ucinput"
	"demo/go-clean-demo/usecase/ucoutput"
)

type LocationUseCase struct {
}

func InitLocationUseCase() *LocationUseCase {
	return &LocationUseCase{}
}

func (uc *LocationUseCase) AddLocation(inputData ucinput.NewLocation, pSuccess *presenter.LocationPresenter) {
	db := fakedb.InitDBConnection()

	locationModel := model.InitLocationModel(db);
	locationModel.AddNewLocation(entity.Location{
		Time: inputData.Time,
		Lat:  inputData.Lat,
		Long: inputData.Long,
	})

	locationList := locationModel.GetAll()

	output := []ucoutput.Location{}

	for _, v := range locationList {
		item := ucoutput.Location{
			Id:    *v.Id,
			Time:  v.Time,
			Lat:   v.Lat,
			Long:  v.Long,
			Speed: int(v.GetSpeedMPH()),
		}

		output = append(output, item)
	}

	pSuccess.PresentAddLocationResponse(output)
}
