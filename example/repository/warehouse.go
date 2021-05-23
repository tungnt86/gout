package repository

import (
	"context"
	"database/sql"

	"github.com/tungnt/gout/example/model"
)

/*
Generate mock struct for unit test
go get github.com/stretchr/testify/mock
go get github.com/vektra/mockery/.../
cd ./example/test && mockery --dir ../repository --name WarehouseRepo
*/
type WarehouseRepo interface {
	Create(ctx context.Context, name string, cityID int64) (*model.Warehouse, error)
}

type warehouseRepo struct {
	db *sql.DB
}

func NewWarehouseRepo(db *sql.DB) WarehouseRepo {
	return &warehouseRepo{db: db}
}

func (r *warehouseRepo) Create(ctx context.Context, name string, cityID int64) (*model.Warehouse, error) {
	warehouse := &model.Warehouse{Name: name, CityID: cityID}
	stmt, err := r.db.Prepare("INSERT INTO warehouse(name, city_id) values(?, ?)")
	if err != nil {
		return nil, err
	}
	res, err := stmt.Exec(warehouse.Name, warehouse.CityID)
	if err != nil {
		return nil, err
	}
	warehouse.ID, err = res.LastInsertId()
	if err != nil {
		return nil, err
	}

	return warehouse, nil
}
