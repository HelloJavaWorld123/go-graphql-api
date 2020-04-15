package postgres

import (
	"database/sql"
	"errors"
	"fmt"
	"time"
)

type Db struct {
	DB *sql.DB
}

func NewConnection(connectionParam string) (*Db, error) {
	db, err := sql.Open("postgres", connectionParam)
	if err != nil {
		return nil, errors.New(err.Error())
	}
	//上面声明的变量二次使用 不用重新简单声明 否则就是一个新的变量
	err = db.Ping()
	if err != nil {
		return nil, errors.New(err.Error())
	}

	db.SetMaxOpenConns(20)
	db.SetMaxIdleConns(10)
	db.SetConnMaxLifetime(time.Duration(time.Duration.Minutes(5)))

	return &Db{DB: db}, nil
}

func ConnectParam(host string, port int64, userName string, dbName string) (string, error) {
	if host == "" || userName == "" || dbName == "" || port < 0 {
		return "", errors.New("缺少数据库连接的参数")
	}
	return fmt.Sprintf("host=%s port=%d username=%s dbname=%s sslmode=disable", host, port, userName, dbName), nil
}
