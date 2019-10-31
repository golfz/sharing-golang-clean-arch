package fakedb

import "time"

type Location struct {
	Id       *int64
	Time     time.Time
	Lat      float64
	Long     float64
	DeviceId int64
}

var db *LocationCollection

type LocationCollection struct {
	collection []Location
}

func InitLocationCollection() *LocationCollection {
	if db == nil {
		db = &LocationCollection{}
	}
	return db
}

func (d *LocationCollection) AddNewLocation(location Location) {
	newId := int64(len(d.collection) + 1)
	location.Id = &newId
	d.collection = append(d.collection, location)
}

func (d *LocationCollection) GetAll() []Location {
	return d.collection
}
