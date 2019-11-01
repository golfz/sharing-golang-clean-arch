package fakedb

import (
	"math/rand"
	"time"
)

type Location struct {
	Id       *int64
	Time     time.Time
	Lat      float64
	Long     float64
	DeviceId int64
}

func (d *Location) GetSpeedMPH() int64 {
	return int64(rand.Intn(40) + 60) // 60 - 100
}

////////////////////////////////////////////////////////////

var db *DBConnection

type DBConnection struct {
	locationTable []Location
}

func InitDBConnection() *DBConnection {
	if db == nil {
		db = &DBConnection{}
	}
	return db
}

/////////////////////////////////////////////////////////////

type LocationModel struct {
	db *DBConnection
}

func InitLocationModel(db *DBConnection) *LocationModel {
	return &LocationModel{db: db}
}

func (d *LocationModel) AddNewLocation(location Location) {
	newId := int64(len(d.db.locationTable) + 1)
	location.Id = &newId
	d.db.locationTable = append(d.db.locationTable, location)
}

func (d *LocationModel) GetAll() []Location {
	return d.db.locationTable
}
