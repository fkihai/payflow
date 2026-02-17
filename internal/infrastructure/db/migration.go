package db

type Migrations interface {
	MigrateUp() error
	MigrateDown() error
}
