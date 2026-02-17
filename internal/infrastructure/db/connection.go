package db

import "database/sql"

type Connection interface {
	Connect() (*sql.DB, error)
}
