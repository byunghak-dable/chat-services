package mysql

import (
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Mysql struct {
	logger *log.Logger
	*gorm.DB
}

func New(logger *log.Logger, user, password, host, port, database string) (*Mysql, error) {
	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN: fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", user, password, host, port, database),
	}), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return &Mysql{logger: logger, DB: db}, nil
}

func (sql *Mysql) Close() {
	db, err := sql.DB.DB()
	if err != nil {
		sql.logger.Printf("get mysql db failure: %v", err)
		return
	}
	err = db.Close()
	if err != nil {
		sql.logger.Printf("close mysql failure: %v", err)
		return
	}
	sql.logger.Println("mysql successfully closed")
}
