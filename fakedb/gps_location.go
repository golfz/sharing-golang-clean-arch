package fakedb

import "time"

type Location struct {
	id int64
	time time.Time
	lat  float64
	long float64
}

type LocationCollection struct {
	locationCollection []Location
}

func InitLocationCollection() *LocationCollection {
	return &LocationCollection{}
}

func (d *LocationCollection) AddNewLocation(location Location) {
	d.locationCollection = append(d.locationCollection, location)
}
