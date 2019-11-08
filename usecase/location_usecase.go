package usecase

import (
	"demo/go-clean-demo/dao"
	"demo/go-clean-demo/entity"
	"demo/go-clean-demo/presenter"
	"demo/go-clean-demo/presenter/viewmodel"
	"demo/go-clean-demo/usecase/ucinput"
	"demo/go-clean-demo/usecase/ucoutput"
	"net/http"
)

type LocationUseCase struct {
	daoFactory *dao.DaoFactory
	pError     *presenter.ErrorPresenter
}

func InitLocationUseCase(daoFactory *dao.DaoFactory, pError *presenter.ErrorPresenter) *LocationUseCase {
	return &LocationUseCase{
		daoFactory: daoFactory,
		pError:     pError,
	}
}

func (uc *LocationUseCase) AddLocation(inputData ucinput.NewLocation, pSuccess *presenter.LocationPresenter) {

	locationDao := uc.daoFactory.GetLocationDao();
	errAdd := locationDao.AddNewLocation(entity.Location{
		Time: inputData.Time,
		Lat:  inputData.Lat,
		Long: inputData.Long,
	})
	if errAdd != nil {
		uc.pError.PresentErrorResponse(http.StatusInternalServerError, viewmodel.ErrorMessage{
			ErrorCode: http.StatusInternalServerError,
			ErrorMsg: "Can not add a new location",
		})
	}

	locationList, errGet := locationDao.GetAll()
	if errGet != nil {
		uc.pError.PresentErrorResponse(http.StatusInternalServerError, viewmodel.ErrorMessage{
			ErrorCode: http.StatusInternalServerError,
			ErrorMsg: "Can not add a new location",
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
