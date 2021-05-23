package gout

import (
	"database/sql"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/suite"

	"github.com/tungnt/gout/database"
)

type UT struct {
	suite.Suite
	DBmock  *sql.DB
	SQLmock sqlmock.Sqlmock
}

func (u *UT) SetupTest() {
	u.DBmock, u.SQLmock = database.NewProvider().MockDB()
}

func (u *UT) TearDownTest() {
	if u.DBmock != nil {
		_ = u.DBmock.Close()
	}
}
