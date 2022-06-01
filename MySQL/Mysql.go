package mysql

import (
	"database/sql"
	"fmt"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file" // required
)

func InitMysqlDB(login, pass, host, port, dbName, args, pathToMigrationFiles string) (*sql.DB, error) {
	dbString := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?%s",
		login, pass, host, port, dbName, args,
	)

	db, err := sql.Open("mysql", dbString)
	if err != nil {
		return nil, err
	}

	err = runMysqlMigrations(db, pathToMigrationFiles, dbName)
	if err != nil {
		return nil, err
	}

	return db, nil
}

func runMysqlMigrations(db *sql.DB, path, dbName string) error {
	driver, err := mysql.WithInstance(db, &mysql.Config{})
	if err != nil {
		return err
	}

	m, err := migrate.NewWithDatabaseInstance(
		path, //file://migrations
		dbName,
		driver)
	if err != nil {
		return err
	}

	err = m.Up()
	if err != nil && err != migrate.ErrNoChange && err != migrate.ErrNilVersion {
		return err
	}

	return nil
}
