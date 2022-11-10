package mysql

import (
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Mysql struct {
	logger *log.Logger
	*gorm.DB
}

func New(logger *log.Logger, dsn string) *Mysql {
	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN: dsn,
	}), &gorm.Config{})
	if err != nil {
		logger.Fatalf("open mysql failed: %s", err)
	}
	return &Mysql{
		logger: logger,
		DB:     db,
	}
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
