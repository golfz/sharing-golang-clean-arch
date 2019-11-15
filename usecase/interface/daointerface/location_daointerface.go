package daointerface

import "demo/go-clean-demo/entity"

type LocationAdder interface {
	AddNewLocation(location entity.Location) error
}

type AllLocationGetter interface {
	GetAll() ([]entity.Location, error)
}
