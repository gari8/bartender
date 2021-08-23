package bartender

import (
	"database/sql"
	"database/sql/driver"
	"fmt"
	"sync"
)

const (
	Postgres = "postgres"
	MySQL = "mysql"
)

var (
	_mux sync.RWMutex
	drivers = make(map[string]driver.Driver)
)

type dsnConnector struct {
	dsn    string
	driver driver.Driver
}


type Cocktail struct {

}

func NewCocktail(driverName string, source string) (*Cocktail, error) {
	_mux.RLock()
	driveri, ok := drivers[driverName]
	_mux.RUnlock()
	if !ok {
		return nil, fmt.Errorf("sql: unknown driver %q (forgotten import?)", driverName)
	}

	if driverCtx, ok := driveri.(driver.DriverContext); ok {
		connector, err := driverCtx.OpenConnector(source)
		if err != nil {
			return nil, err
		}
		return OpenDB(connector), nil
	}

	return OpenDB(dsnConnector{dsn: source, driver: driveri}), nil
}
