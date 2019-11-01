package model

import (
	"demo/go-clean-demo/entity"
	"demo/go-clean-demo/fakedb"
)

type LocationModel struct {
	db *fakedb.DBConnection
}

func InitLocationModel(db *fakedb.DBConnection) *LocationModel {
	return &LocationModel{db: db}
}

func (d *LocationModel) AddNewLocation(location entity.Location) {
	newId := int64(len(d.db.LocationTable) + 1)
	location.Id = &newId
	d.db.LocationTable = append(d.db.LocationTable, location)
}

func (d *LocationModel) GetAll() []entity.Location {
	return d.db.LocationTable
}
