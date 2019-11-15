package daointerface

type DaoFactory interface {
	GetLocationAdder() LocationAdder
	GetAllLocationGetter() AllLocationGetter
}
