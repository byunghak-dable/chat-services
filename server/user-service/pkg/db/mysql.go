package db

import (
	"fmt"

	"github.com/widcraft/user-service/pkg/logger"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Mysql struct {
	logger logger.Logger
	*gorm.DB
}

func NewMysql(logger logger.Logger, user, password, host, port, database string) (*Mysql, error) {
	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN: fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", user, password, host, port, database),
	}), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return &Mysql{logger: logger, DB: db}, nil
}

func (sql *Mysql) Close() error {
	db, err := sql.DB.DB()
	if err != nil {
		return err
	}
	return db.Close()
}
