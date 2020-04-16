package postgres

import (
	"fmt"
	"log"
	"testing"
)

func TestConnectParam(t *testing.T) {
	host := "172.16.20.190"
	var port int64 = 3306
	userName := "admin"
	password := "admin"
	dbName := "go-postgresql-api"

	connectParam, err := ConnectParam(host, port, userName, password, dbName)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(connectParam)
}

func TestNewConnection(t *testing.T) {
	host := "172.16.20.195"
	var port int64 = 5432
	userName := "postgres"
	password := "postgres"
	dbName := "postgres"

	param, err := ConnectParam(host, port, userName, password, dbName)
	if err != nil {
		t.Fatal(err)
	}

	connection, err := NewConnection(param)
	if err != nil {
		t.Fatal(err)
	}
	err = connection.Ping()
	if err != nil {
		t.Fatal(err)
	}
	t.Log("链接成功")

}
