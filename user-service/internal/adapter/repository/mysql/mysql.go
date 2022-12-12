package mysql

import (
	"fmt"

	log "github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Mysql struct {
	logger log.FieldLogger
	*gorm.DB
}

func New(logger log.FieldLogger, user, password, host, port, database string) (*Mysql, error) {
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
