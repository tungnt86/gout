package usecase

import (
	"context"
	"errors"

	"github.com/tungnt/gout"
	"github.com/tungnt/gout/example/model"
	"github.com/tungnt/gout/example/test/mocks"
)

type CreateWarehoseTestSuite struct {
	gout.UT
	cityRepoMock      *mocks.CityRepo
	warehouseRepoMock *mocks.WarehouseRepo
	warehouseUsecase  CreateWarehouse
}

func (s *CreateWarehoseTestSuite) SetupTest() {
	s.UT.SetupTest()
	s.cityRepoMock = &mocks.CityRepo{}
	s.warehouseRepoMock = &mocks.WarehouseRepo{}
	s.warehouseUsecase = NewCreateWarehouse(
		s.DBmock,
		s.cityRepoMock,
		s.warehouseRepoMock,
	)
}

func (s *CreateWarehoseTestSuite) TestCreate_WarehoueError_Rollback() {
	cityName := "City Name"
	whName := "WH Name"
	s.SQLmock.ExpectBegin()
	s.SQLmock.ExpectRollback()
	s.cityRepoMock.On("Create", context.Background(), cityName).Return(&model.City{ID: int64(1)}, nil).Once()
	s.warehouseRepoMock.On("Create", context.Background(), whName, int64(1)).Return(nil, errors.New("Any error")).Once()
	_, err := s.warehouseUsecase.CreateWarehouseAndCity(context.Background(), whName, cityName)
	s.Error(err)
}
