package mysql

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitMysqlGormDB(login, pass, host, port, dbName, args string) (*gorm.DB, error) {
	dbString := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?%s",
		login, pass, host, port, dbName, args,
	)

	db, err := gorm.Open(mysql.Open(dbString), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return db, nil
}
