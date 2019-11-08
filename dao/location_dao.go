package dao

import (
	"demo/go-clean-demo/entity"
	"demo/go-clean-demo/fakedb"
)

type LocationDao struct {
	db *fakedb.DBConnection
}

func InitLocationDao(db *fakedb.DBConnection) *LocationDao {
	return &LocationDao{db: db}
}

func (d *LocationDao) AddNewLocation(location entity.Location) error {
	newId := int64(len(d.db.LocationTable) + 1)
	location.Id = &newId
	d.db.LocationTable = append(d.db.LocationTable, location)
	return nil
}

func (d *LocationDao) GetAll() ([]entity.Location, error) {
	return d.db.LocationTable, nil
}
