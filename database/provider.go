package database

import (
	"database/sql"
	"sync"

	"github.com/DATA-DOG/go-sqlmock"

	"github.com/tungnt/gout/must"
)

var (
	once     sync.Once
	mutex    = &sync.Mutex{}
	instance *provider
)

type Provider interface {
	MockDB() (*sql.DB, sqlmock.Sqlmock)
}

type provider struct{}

func NewProvider() Provider {
	once.Do(func() {
		instance = &provider{}
	})

	return instance
}

func (p *provider) MockDB() (*sql.DB, sqlmock.Sqlmock) {
	mutex.Lock()
	defer mutex.Unlock()

	mockDB, sqlMock, err := sqlmock.New()
	must.NotFail(err)
	return mockDB, sqlMock
}
