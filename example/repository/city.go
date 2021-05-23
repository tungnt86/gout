package repository

import (
	"context"
	"database/sql"

	"github.com/tungnt/gout/example/model"
)

/*
Generate mock struct for unit test
go get github.com/stretchr/testify/mock
go get github.com/vektra/mockery/v2/.../
cd example/test && mockery --dir ../repository --name CityRepo
*/
type CityRepo interface {
	Create(ctx context.Context, name string) (*model.City, error)
}

type cityRepo struct {
	db *sql.DB
}

func NewCityRepo(db *sql.DB) CityRepo {
	return &cityRepo{db: db}
}

func (r *cityRepo) Create(ctx context.Context, name string) (*model.City, error) {
	city := &model.City{Name: name}
	stmt, err := r.db.Prepare("INSERT INTO city(name) values(?)")
	if err != nil {
		return nil, err
	}
	res, err := stmt.Exec(city.Name)
	if err != nil {
		return nil, err
	}
	city.ID, err = res.LastInsertId()
	if err != nil {
		return nil, err
	}

	return city, nil
}
