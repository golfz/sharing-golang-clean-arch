package usecase

import (
	"demo/go-clean-demo/entity"
	"demo/go-clean-demo/usecase/interface/daointerface"
	"demo/go-clean-demo/usecase/interface/pinterface"
	"demo/go-clean-demo/usecase/ucinput"
	"demo/go-clean-demo/usecase/ucoutput"
	"net/http"
)

type LocationUseCase struct {
	daoFactory daointerface.DaoFactory
	pError     pinterface.ErrorResponsePresenter
}

func InitLocationUseCase(daoFactory daointerface.DaoFactory, pError pinterface.ErrorResponsePresenter) *LocationUseCase {
	return &LocationUseCase{
		daoFactory: daoFactory,
		pError:     pError,
	}
}

func (uc *LocationUseCase) AddLocation(inputData ucinput.NewLocation, pSuccess pinterface.AddLocationResponsePresenter) {

	locationAdder := uc.daoFactory.GetLocationAdder()

	errAdd := locationAdder.AddNewLocation(entity.Location{
		Time: inputData.Time,
		Lat:  inputData.Lat,
		Long: inputData.Long,
	})
	if errAdd != nil {
		uc.pError.PresentErrorResponse(ucoutput.Error{
			ErrorStatus:  http.StatusInternalServerError,
			ErrorCode:    http.StatusInternalServerError,
			ErrorMessage: "Can not add a new location",
		})
	}

	locationGetter := uc.daoFactory.GetAllLocationGetter()

	locationList, errGet := locationGetter.GetAll()
	if errGet != nil {
		uc.pError.PresentErrorResponse(ucoutput.Error{
			ErrorStatus:  http.StatusInternalServerError,
			ErrorCode:    http.StatusInternalServerError,
			ErrorMessage: "Can not get all location after added",
		})
	}

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
