package db

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

type MysqlConfig struct {
	User     string
	Password string
	Database string
	Host     string
	Port     string
}

type Mysql struct {
	*sql.DB
}

func NewMysql(config MysqlConfig) (*Mysql, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", config.User, config.Password, config.Host, config.Port, config.Database)
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal(err)
	}

	return &Mysql{DB: db}, nil
}
