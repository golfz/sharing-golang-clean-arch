package dao

import (
	"demo/go-clean-demo/fakedb"
	"demo/go-clean-demo/usecase/interface/daointerface"
)

type DaoFactory struct {
	db *fakedb.DBConnection
}

func InitDaoFactory(db *fakedb.DBConnection) *DaoFactory {
	return &DaoFactory{
		db: db,
	}
}

func (d *DaoFactory) GetLocationAdder() daointerface.LocationAdder {
	return &LocationDao{db: d.db}
}

func (d *DaoFactory) GetAllLocationGetter() daointerface.AllLocationGetter {
	return &LocationDao{db: d.db}
}
