package usecase

import (
	"demo/go-clean-demo/entity"
	"demo/go-clean-demo/fakedb"
	"demo/go-clean-demo/model"
	"demo/go-clean-demo/presenter"
	"demo/go-clean-demo/usecase/ucinput"
)

type LocationUseCase struct {
}

func InitLocationUseCase() *LocationUseCase {
	return &LocationUseCase{}
}

func (uc *LocationUseCase) AddLocation(inputData ucinput.NewLocation, pSuccess presenter.LocationPresenter) {
	db := fakedb.InitDBConnection()

	locationModel := model.InitLocationModel(db);
	locationModel.AddNewLocation(entity.Location{
		Time: inputData.Time,
		Lat:  inputData.Lat,
		Long: inputData.Long,
	})

	locationList := locationModel.GetAll()

	pSuccess.PresentAddLocationResponse(locationList)
}
