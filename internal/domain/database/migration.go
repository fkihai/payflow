package database

type Migrations interface {
	MigrateUp() error
	MigrateDown() error
}
