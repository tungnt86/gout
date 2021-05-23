package usecase

import (
	"context"
	"database/sql"

	"github.com/tungnt/gout/example/model"
	"github.com/tungnt/gout/example/repository"
)

type CreateWarehouse interface {
	CreateWarehouseAndCity(ctx context.Context, warehouseName, cityName string) (*model.Warehouse, error)
}
type createWarehouse struct {
	db            *sql.DB
	cityRepo      repository.CityRepo
	warehouseRepo repository.WarehouseRepo
}

func NewCreateWarehouse(
	db *sql.DB,
	cityRepo repository.CityRepo,
	warehouseRepo repository.WarehouseRepo,
) *createWarehouse {
	return &createWarehouse{
		db:            db,
		cityRepo:      cityRepo,
		warehouseRepo: warehouseRepo,
	}
}

func (w *createWarehouse) CreateWarehouseAndCity(ctx context.Context, warehouseName, cityName string) (*model.Warehouse, error) {
	tx, err := w.db.BeginTx(ctx, &sql.TxOptions{Isolation: sql.LevelSerializable})
	if err != nil {
		return nil, err
	}
	city, err := w.cityRepo.Create(ctx, cityName)
	if err != nil {
		_ = tx.Rollback()
		return nil, err
	}
	warehouse, err := w.warehouseRepo.Create(ctx, warehouseName, city.ID)
	if err != nil {
		_ = tx.Rollback()
		return nil, err
	}
	err = tx.Commit()
	if nil != err {
		return nil, err
	}
	return warehouse, nil
}
