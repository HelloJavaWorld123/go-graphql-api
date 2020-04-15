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
	dbName := "go-postgresql-api"

	connectParam, err := ConnectParam(host, port, userName, dbName)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(connectParam)
}
