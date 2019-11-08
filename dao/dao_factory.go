package dao

import "demo/go-clean-demo/fakedb"

type DaoFactory struct {
	db *fakedb.DBConnection
}

func InitDaoFactory(db *fakedb.DBConnection) *DaoFactory {
	return &DaoFactory{
		db: db,
	}
}

func (d *DaoFactory) GetLocationDao() *LocationDao {
	return &LocationDao{db: d.db}
}
