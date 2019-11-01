package fakedb

import "demo/go-clean-demo/entity"

var db *DBConnection

type DBConnection struct {
	LocationTable []entity.Location
}

func InitDBConnection() *DBConnection {
	if db == nil {
		db = &DBConnection{}
	}
	return db
}

/////////////////////////////////////////////////////////////
