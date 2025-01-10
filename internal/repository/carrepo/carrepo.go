package carrepo

import (
	_ "database/sql"

	_ "github.com/go-pg/pg/v10"
	_ "github.com/go-pg/pg/v10/orm"
)

type Car struct {
	ID            int
	somethingElse string
	Type          int
}

func (c Car) Validate() (bool, error) {
	return c.ID > 0, nil
}
